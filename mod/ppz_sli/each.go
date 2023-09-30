package ppz_sli

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
