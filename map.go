package channels

import "context"

func Map[T any, TMap any](ctx context.Context, source Stream[T], f func(ctx context.Context, element T) TMap) Stream[TMap] {
	out := make(chan TMap)
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
				out <- f(ctx, element)
			}
		}
	}()
	return Stream[TMap]{channel: out}
}
