package channels

import "context"

func Filter[T any](ctx context.Context, source <-chan T, predicate func(ctx context.Context, element T) bool) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case element, ok := <-source:
				if !ok {
					return
				}
				if predicate(ctx, element) {
					out <- element
				}
			}
		}
	}()
	return out
}
