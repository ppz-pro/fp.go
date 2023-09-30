package ppz_sli

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
