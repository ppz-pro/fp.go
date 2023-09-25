package slice

func make_container[item any](source []item) []item {
	return make([]item, 0, len(source))
}

type Item[El any] struct {
	el    El
	index int
	slice Slice[El]
}
