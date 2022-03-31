package channels

import "context"

func Map[T any, TMap any](ctx context.Context, source <-chan T, f func(ctx context.Context, element T) TMap) <-chan TMap {
	out := make(chan TMap)
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
				out <- f(ctx, element)
			}
		}
	}()
	return out
}
