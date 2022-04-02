package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := FromSlice([]int{1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	subject := Sum(numbers)
	assert.Equal(t, subject, 58)
}

func TestSumWhenEmpty(t *testing.T) {
	numbers := Of[int]()

	subject := Sum(numbers)
	assert.Equal(t, subject, 0)
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numbers := FromSlice(rangeInt(1, 10))

		Sum(numbers)
	}
}
