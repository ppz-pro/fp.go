package ppz_sli

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

func make_container[item any](size int) []item {
	return make([]item, size)
}

func (slice Slice[El]) Each3(cb func(El, int, Slice[El])) {
	for index, el := range slice {
		cb(el, index, slice)
	}
}
func (slice Slice[El]) Each2(cb func(El, int)) {
	slice.Each3(func(el El, index int, _ Slice[El]) {
		cb(el, index)
	})
}
func (slice Slice[El]) Each(cb func(El)) {
	slice.Each3(func(el El, _ int, _ Slice[El]) {
		cb(el)
	})
}

func (slice Slice[El]) Filter3(cb func(El, int, Slice[El]) bool) Slice[El] {
	var result Slice[El]
	for index, item := range slice {
		if cb(item, index, slice) {
			result = append(result, item)
		}
	}
	return result
}
func (slice Slice[El]) Filter2(cb func(El, int) bool) Slice[El] {
	return slice.Filter3(func(el El, index int, _ Slice[El]) bool {
		return cb(el, index)
	})
}
func (slice Slice[El]) Filter(cb func(El) bool) Slice[El] {
	return slice.Filter3(func(el El, index int, _ Slice[El]) bool {
		return cb(el)
	})
}

func (slice Slice[El]) Find3(cb func(El, int, Slice[El]) bool) (El, int, bool) {
	for index, item := range slice {
		if cb(item, index, slice) {
			return item, index, true
		}
	}
	var nil_item El
	return nil_item, 0, false
}
func (slice Slice[El]) Find2(cb func(El, int) bool) (El, int, bool) {
	return slice.Find3(func(el El, index int, _ Slice[El]) bool {
		return cb(el, index)
	})
}
func (slice Slice[El]) Find(cb func(El) bool) (El, int, bool) {
	return slice.Find3(func(el El, _ int, _ Slice[El]) bool {
		return cb(el)
	})
}

func Map3[Input any, Output any](inputs []Input, cb func(Input, int, []Input) Output) Slice[Output] {
	result := make_container[Output](len(inputs))
	for index, item := range inputs {
		result[index] = cb(item, index, inputs)
	}
	return result
}
func Map2[Input any, Output any](inputs []Input, cb func(Input, int) Output) Slice[Output] {
	return Map3(inputs, func(el Input, index int, _ []Input) Output {
		return cb(el, index)
	})
}
func Map[Input any, Output any](inputs []Input, cb func(Input) Output) Slice[Output] {
	return Map3(inputs, func(el Input, _ int, _ []Input) Output {
		return cb(el)
	})
}

func Reduce3[El any, Result any](
	result *Result,
	slice []El,
	cb func(*Result, El, int, []El),
) Result {
	for index, el := range slice {
		cb(result, el, index, slice)
	}
	return *result
}
func Reduce2[El any, Result any](
	result *Result,
	slice []El,
	cb func(*Result, El, int),
) Result {
	return Reduce3[El, Result](
		result,
		slice,
		func(result *Result, el El, index int, _ []El) {
			cb(result, el, index)
		},
	)
}
func Reduce[El any, Result any](
	result *Result,
	slice []El,
	cb func(*Result, El),
) Result {
	return Reduce3[El, Result](
		result,
		slice,
		func(result *Result, el El, _ int, _ []El) {
			cb(result, el)
		},
	)
}

func (slice Slice[El]) Reverse() Slice[El] {
	length := len(slice)
	result := make_container[El](length)
	for i, item := range slice {
		result[length-1-i] = item
	}
	return result
}

func (slice Slice[El]) Every3(cb func(El, int, Slice[El]) bool) bool {
	for index, item := range slice {
		if !cb(item, index, slice) {
			return false
		}
	}
	return true
}
func (slice Slice[El]) Every2(cb func(El, int) bool) bool {
	return slice.Every3(
		func(el El, index int, _ Slice[El]) bool {
			return cb(el, index)
		},
	)
}
func (slice Slice[El]) Every(cb func(El) bool) bool {
	return slice.Every3(
		func(el El, _ int, _ Slice[El]) bool {
			return cb(el)
		},
	)
}

func (slice Slice[El]) Some3(cb func(El, int, Slice[El]) bool) bool {
	for index, item := range slice {
		if cb(item, index, slice) {
			return true
		}
	}
	return false
}
func (slice Slice[El]) Some2(cb func(El, int) bool) bool {
	return slice.Some3(
		func(el El, index int, _ Slice[El]) bool {
			return cb(el, index)
		},
	)
}
func (slice Slice[El]) Some(cb func(El) bool) bool {
	return slice.Some3(
		func(el El, _ int, _ Slice[El]) bool {
			return cb(el)
		},
	)
}
