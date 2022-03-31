package channels

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	numbers := FromSlice(rangeInt(1, 5))
	result := Map(context.TODO(), numbers, func(ctx context.Context, element int) int { return element * 2 })
	subject := ToSlice(result)
	assert.Len(t, subject, 4)
	assert.ElementsMatch(t, []int{2, 4, 6, 8}, subject)
}

func BenchmarkMap(b *testing.B) {
	ctx := context.TODO()
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(rangeInt(1, 10))
		ToSlice(Map(ctx, numbers, func(ctx context.Context, element int) bool { return element%2 == 0 }))
	}
}
