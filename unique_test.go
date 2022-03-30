package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	numbers := make(chan int, 10)
	go func() {
		for _, a := range []int{1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
			numbers <- a
		}
		close(numbers)
	}()

	result := Unique(numbers, func(element int) int { return element })
	subject := ToSlice(result)
	assert.Len(t, subject, 10)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, subject)
}

func BenchmarkUnique(b *testing.B) {
	for n := 0; n < b.N; n++ {
		numbers := make(chan int, 10)
		go func() {
			for _, a := range []int{1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
				numbers <- a
			}
			close(numbers)
		}()

		result := Unique(numbers, func(element int) int { return element })
		ToSlice(result)
	}
}
