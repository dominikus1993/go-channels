package channels

func ToSlice[T any](s Stream[T]) []T {
	res := make([]T, 0)
	for v := range s.channel {
		res = append(res, v)
	}
	return res
}

func FromSlice[T any](s []T) Stream[T] {
	res := make(chan T)
	go func() {
		for _, v := range s {
			res <- v
		}
		close(res)
	}()
	return Stream[T]{channel: res}
}

func Of[T any](s ...T) Stream[T] {
	res := make(chan T)
	go func() {
		for _, v := range s {
			res <- v
		}
		close(res)
	}()
	return Stream[T]{channel: res}
}
