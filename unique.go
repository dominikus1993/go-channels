package channels

func Unique[T any, TID comparable](stream Stream[T], f func(elem T) TID) Stream[T] {
	seen := make(map[TID]bool)
	res := make(chan T, 20)
	go func() {
		for v := range stream.channel {
			id := f(v)
			if !seen[id] {
				seen[id] = true
				res <- v
			}
		}
		close(res)
	}()
	return Stream[T]{channel: res}
}

func UniqueComparable[T comparable](stream Stream[T]) Stream[T] {
	seen := make(map[T]bool)
	res := make(chan T, 20)
	go func() {
		for v := range stream.channel {
			if !seen[v] {
				seen[v] = true
				res <- v
			}
		}
		close(res)
	}()
	return Stream[T]{channel: res}
}
