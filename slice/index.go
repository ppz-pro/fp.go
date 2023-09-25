package slice

type Slice[El any] []El

func New[El any](el []El) Slice[El] {
	return Slice[El](el)
}

func (slice Slice[El]) Raw() []El {
	return []El(slice)
}

type Item[El any] struct {
	El    El
	Index int
	Slice Slice[El]
}

func make_container[item any](size int) []item {
	return make([]item, 0, size)
}

func (slice Slice[El]) Each(cb func(item Item[El])) {
	for index, item := range slice {
		cb(Item[El]{item, index, slice})
	}
}

func (slice Slice[El]) Filter(cb func(item Item[El]) bool) Slice[El] {
	var result []El
	for index, item := range slice {
		if cb(Item[El]{item, index, slice}) {
			result = append(result, item)
		}
	}
	return New[El](result)
}
