package ppz_sli

import "testing"

func is_same_slice[El comparable](a []El, b []El) bool {
	for i, el := range a {
		if el != b[i] {
			return false
		}
	}
	return len(a) == len(b) && cap(a) == cap(b)
}

func Test_is_same_slice(t *testing.T) {
	if is_same_slice[int](
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	) == false {
		t.Error("same slice not work")
	}

	if is_same_slice[int](
		[]int{1, 2, 3},
		[]int{1, 2, 4},
	) {
		t.Error("same slice not work")
	}
}
