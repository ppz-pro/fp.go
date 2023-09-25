package slice

type Slice[El any] []El

func Arm[El any](el []El) Slice[El] {
	return Slice[El](el)
}

func New[El any](el ...El) Slice[El] {
	return Slice[El](el)
}

func (slice Slice[El]) Disarm() []El {
	return []El(slice)
}

type Item[El any] struct {
	El    El
	Index int
	Slice Slice[El]
}

func make_container[item any](size int) []item {
	return make([]item, size)
}

func (slice Slice[El]) Each(cb func(item Item[El])) {
	for index, item := range slice {
		cb(Item[El]{item, index, slice})
	}
}

func (slice Slice[El]) Filter(cb func(item Item[El]) bool) Slice[El] {
	var result Slice[El]
	for index, item := range slice {
		if cb(Item[El]{item, index, slice}) {
			result = append(result, item)
		}
	}
	return result
}

func (slice Slice[El]) Find(cb func(item Item[El]) bool) (El, int, bool) {
	for index, item := range slice {
		if cb(Item[El]{item, index, slice}) {
			return item, index, true
		}
	}
	var nil_item El
	return nil_item, 0, false
}

func Map[Input any, Output any](inputs []Input, cb func(Item[Input]) Output) Slice[Output] {
	result := make_container[Output](len(inputs))
	for index, item := range inputs {
		result[index] = cb(Item[Input]{item, index, inputs})
	}
	return result
}

func (slice Slice[El]) Reverse() Slice[El] {
	length := len(slice)
	result := make_container[El](length)
	for i, item := range slice {
		result[length-1-i] = item
	}
	return result
}

func (slice Slice[El]) Every(cb func(Item[El]) bool) bool {
	for index, item := range slice {
		if !cb(Item[El]{item, index, slice}) {
			return false
		}
	}
	return true
}

func (slice Slice[El]) Some(cb func(Item[El]) bool) bool {
	for index, item := range slice {
		if cb(Item[El]{item, index, slice}) {
			return true
		}
	}
	return false
}
