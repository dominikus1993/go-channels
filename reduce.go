package channels

func Reduce[T any, TRes any](s Stream[T], zero TRes, f func(acc TRes, elem T) TRes) TRes {
	res := zero
	for v := range s.channel {
		res = f(res, v)
	}
	return res
}
