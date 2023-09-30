package ppz_sli

import "testing"

func Test_filter(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5, 6, 7}
	odds := Arm(raw).Filter(
		func(el int) bool {
			return el%2 == 1
		},
	).Disarm()

	if len(odds) != 4 {
		t.Error("filter error, length", len(odds))
	}
	if !is_same_slice[int](odds, []int{1, 3, 5, 7}) {
		t.Error("filter error, elements", odds)
	}
}
