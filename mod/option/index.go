package option

type Option[V any] struct {
	value     V
	has_value bool
}

func New[V any](v V) Option[V] {
	return Option[V]{v, true}
}

func Nil[V any]() Option[V] {
	return Option[V]{has_value: false}
}

func (opt Option[V]) Match(no_value func(), has_value func(V)) {
	if opt.has_value {
		has_value(opt.value)
	} else {
		no_value()
	}
}

func Match[V any, Result any](
	opt Option[V],
	no_value func() Result,
	has_value func(V) Result,
) Result {
	if opt.has_value {
		return has_value(opt.value)
	} else {
		return no_value()
	}
}
