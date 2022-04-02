package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceSum(t *testing.T) {
	numbers := FromSlice([]int{1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	subject := Reduce(numbers, 0, func(acc int, elem int) int { return acc + elem })
	assert.Equal(t, subject, 58)
}

func TestReduceSumString(t *testing.T) {
	numbers := FromSlice([]string{"Hello", " ", "World"})

	subject := Reduce(numbers, "", func(acc string, elem string) string { return acc + elem })
	assert.Equal(t, subject, "Hello World")
}

func TestReduceSumWhenEmpty(t *testing.T) {
	numbers := Of[int]()

	subject := Reduce(numbers, 0, func(acc int, elem int) int { return acc + elem })
	assert.Equal(t, subject, 0)
}

func BenchmarkReduceSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(rangeInt(1, 10))

		Reduce(numbers, 0, func(acc int, elem int) int { return acc + elem })
	}
}
