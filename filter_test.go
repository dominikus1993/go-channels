package channels

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	numbers := FromSlice(rangeInt(1, 10))
	result := Filter(numbers, func(element int) bool { return element%2 == 0 })
	subject := ToSlice(result)
	assert.Len(t, subject, 4)
	assert.ElementsMatch(t, []int{2, 4, 6, 8}, subject)
}

func BenchmarkFilter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(rangeInt(1, 10))
		ToSlice(Filter(numbers, func(element int) bool { return element%2 == 0 }))
	}
}

func TestFilterContext(t *testing.T) {
	numbers := FromSlice(rangeInt(1, 10))
	result := FilterContext(context.TODO(), numbers, func(ctx context.Context, element int) bool { return element%2 == 0 })
	subject := ToSlice(result)
	assert.Len(t, subject, 4)
	assert.ElementsMatch(t, []int{2, 4, 6, 8}, subject)
}

func BenchmarkFilterContext(b *testing.B) {
	ctx := context.TODO()
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(rangeInt(1, 10))
		ToSlice(FilterContext(ctx, numbers, func(ctx context.Context, element int) bool { return element%2 == 0 }))
	}
}
