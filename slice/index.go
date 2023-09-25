package slice

type Slice[El any] struct {
	value []El
}

func New[El any](el []El) Slice[El] {
	return Slice[El]{el}
}

type Item[El any] struct {
	el    El
	index int
	slice Slice[El]
}

func make_container[item any](source []item) []item {
	return make([]item, 0, len(source))
}

func (slice Slice[El]) Each(cb func(item Item[El])) {
	for index, item := range slice.value {
		cb(Item[El]{item, index, slice})
	}
}
