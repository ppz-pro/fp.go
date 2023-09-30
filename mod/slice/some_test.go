package ppz_sli

import "testing"

func Test_some(t *testing.T) {
	raw := []int{1, 3, 6, 2, 1, 10, 11}
	count := 0

	has6 := Arm(raw).Some(func(el int) bool {
		count += 1
		return el == 6
	})

	if !has6 {
		t.Error("slice.Some no 6")
	}
	if count != 3 {
		t.Error("slice.Some times of test is not 3", count)
	}
}
