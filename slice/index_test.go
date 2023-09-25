package slice

import "testing"

func is_same_slice[El comparable](a []El, b []El) bool {
	for i, el := range a {
		if el != b[i] {
			return false
		}
	}
	return true
}

func Test_foreach(t *testing.T) {
	raw := []int{1, 2, 3}

	test := make_container[int](raw)
	count := 0

	New(raw).Each(func(item Item[int]) {
		if item.index != count {
			t.Error("index error", item.index, count)
		}
		test = append(test, item.el)
		count = count + 1
	})

	if !is_same_slice(raw, test) {
		t.Error("wrong el after each")
	}
}
