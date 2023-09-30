package ppz_sli

import "testing"

func Test_every(t *testing.T) {
	raw := []int{2, 4, 6, 8, 10}
	count := 0
	is_all_even := Arm(raw).Every(func(el int) bool {
		count += 1
		return el%2 == 0
	})
	if !is_all_even {
		t.Error("slice.every not all even")
	}
	if count != 5 {
		t.Error("slice.every iteration time is not 5", count)
	}
}
