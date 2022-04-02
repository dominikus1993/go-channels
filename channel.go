package channels

func FromChannel[T any](channel chan T) Stream[T] {
	return Stream[T]{channel: channel}
}

func ToReadOnlyChannel[T any](s Stream[T]) <-chan T {
	return s.channel
}
