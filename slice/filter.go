package ppz_sli

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
