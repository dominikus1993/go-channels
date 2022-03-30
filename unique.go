package channels

func Unique[T any, TID comparable](stream <-chan T, f func(elem T) TID) <-chan T {
	seen := make(map[TID]bool)
	res := make(chan T, 20)
	go func() {
		for v := range stream {
			id := f(v)
			if !seen[id] {
				seen[id] = true
				res <- v
			}
		}
		close(res)
	}()
	return res
}

func UniqueComparable[T comparable](stream <-chan T) <-chan T {
	seen := make(map[T]bool)
	res := make(chan T, 20)
	go func() {
		for v := range stream {
			if !seen[v] {
				seen[v] = true
				res <- v
			}
		}
		close(res)
	}()
	return res
}
