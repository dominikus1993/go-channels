package channels

import (
	"context"
	"sync"
)

func FanIn[T any](ctx context.Context, stream ...Stream[T]) Stream[T] {
	var wg sync.WaitGroup
	out := make(chan T)
	output := func(c Stream[T]) {
		defer wg.Done()
		for v := range c.channel {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}
	wg.Add(len(stream))
	for _, c := range stream {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return Stream[T]{channel: out}
}
