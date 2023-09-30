package ppz_sli

import "testing"

func Test_each(t *testing.T) {
	raw := []int{1, 2, 3}

	test := make_container[int](len(raw))
	count := 0

	Arm(raw).Each2(func(el int, index int) {
		if index != count {
			t.Error("index error", index, count)
		}
		test[index] = el
		count = count + 1
	})

	if !is_same_slice([]int{1, 2, 3}, test) {
		t.Error("wrong el after each", test)
	}
}
