package slice

type Slice[El any] struct {
	value []El
}

func New[El any](el []El) Slice[El] {
	return Slice[El]{el}
}

func (slice Slice[El]) Each(cb func(item Item[El])) {
	for index, item := range slice.value {
		cb(Item[El]{item, index, slice})
	}
}
