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

func (opt Option[V]) Match(has_value func(V), no_value func()) {
	if opt.has_value {
		has_value(opt.value)
	} else {
		no_value()
	}
}
