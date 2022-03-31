package channels

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFanIn(t *testing.T) {
	numbers := FromSlice(rangeInt(1, 10))
	numbers2 := FromSlice(rangeInt(10, 20))

	result := FanIn(context.TODO(), numbers, numbers2)
	subject := ToSlice(result)
	assert.Len(t, subject, 19)
	assert.ElementsMatch(t, rangeInt(1, 20), subject)
}

func BenchmarkFanIn(b *testing.B) {
	ctx := context.TODO()
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(rangeInt(1, 10))
		numbers2 := FromSlice(rangeInt(10, 20))

		result := FanIn(ctx, numbers, numbers2)
		ToSlice(result)
	}
}
