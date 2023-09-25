package slice

import (
	"testing"
)

func is_same_slice[El comparable](a []El, b []El) bool {
	for i, el := range a {
		if el != b[i] {
			return false
		}
	}
	return true
}

func Test_each(t *testing.T) {
	raw := []int{1, 2, 3}

	test := make_container[int](len(raw))
	count := 0

	New(raw).Each(func(item Item[int]) {
		if item.Index != count {
			t.Error("index error", item.Index, count)
		}
		test = append(test, item.El)
		count = count + 1
	})

	if !is_same_slice(raw, test) {
		t.Error("wrong el after each")
	}
}

func Test_filter(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5, 6, 7}
	odds := New[int](raw).Filter(
		func(item Item[int]) bool {
			return item.El%2 == 1
		},
	).Get()

	if len(odds) != 4 {
		t.Error("filter error, length", len(odds))
	}
	if !is_same_slice[int](odds, []int{1, 3, 5, 7}) {
		t.Error("filter error, elements", odds)
	}
}
