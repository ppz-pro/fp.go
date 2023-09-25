package slice

func is_same_slice[El comparable](a []El, b []El) bool {
	for i, el := range a {
		if el != b[i] {
			return false
		}
	}
	return true
}
