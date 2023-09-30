package ppz_sli

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
