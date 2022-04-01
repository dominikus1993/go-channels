package channels

type Stream[T any] struct {
	channel <-chan T
}
