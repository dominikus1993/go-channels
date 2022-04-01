package channels

import "context"

func Filter[T any](ctx context.Context, source Stream[T], predicate func(ctx context.Context, element T) bool) Stream[T] {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case element, ok := <-source.channel:
				if !ok {
					return
				}
				if predicate(ctx, element) {
					out <- element
				}
			}
		}
	}()
	return Stream[T]{channel: out}
}
