package channels

type Sumable interface {
	int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

func Sum[T Sumable](stream Stream[T]) T {
	var sum T
	for v := range stream.channel {
		sum += v
	}
	return sum
}
