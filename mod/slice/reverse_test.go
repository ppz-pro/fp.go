package ppz_sli

import "testing"

func Test_reverse(t *testing.T) {
	raw := []int{1, 2, 3, 4, 6}
	reversed := Arm(raw).Reverse()
	if !is_same_slice(reversed, []int{6, 4, 3, 2, 1}) {
		t.Error("reverse error", reversed)
	}
}
