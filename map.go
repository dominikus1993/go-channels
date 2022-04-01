package channels

import "context"

func MapContext[T any, TMap any](ctx context.Context, source Stream[T], f func(ctx context.Context, element T) TMap) Stream[TMap] {
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

func Map[T any, TMap any](source Stream[T], f func(element T) TMap) Stream[TMap] {
	out := make(chan TMap)
	go func() {
		defer close(out)
		for v := range source.channel {
			out <- f(v)
		}
	}()
	return Stream[TMap]{channel: out}
}
