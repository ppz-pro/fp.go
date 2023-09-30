package ppz_sli

func (slice Slice[El]) Reverse() Slice[El] {
	length := len(slice)
	result := make_container[El](length)
	for i, item := range slice {
		result[length-1-i] = item
	}
	return result
}
