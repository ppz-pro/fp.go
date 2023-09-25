package slice

import (
	"fmt"
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

	Arm(raw).Each(func(item Item[int]) {
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
	odds := Arm(raw).Filter(
		func(item Item[int]) bool {
			return item.El%2 == 1
		},
	).Disarm()

	if len(odds) != 4 {
		t.Error("filter error, length", len(odds))
	}
	if !is_same_slice[int](odds, []int{1, 3, 5, 7}) {
		t.Error("filter error, elements", odds)
	}
}

func Test_map(t *testing.T) {
	raw := []int{1, 2, 3}
	power2 := Map(raw, func(item Item[int]) string {
		return fmt.Sprintf("%d", item.El*item.El)
	})

	count := 0

	for index, item := range power2 {
		switch index {
		case 0:
			count = count + 1
			if item != "1" {
				t.Error("map error on first", item)
			}
		case 1:
			count = count + 1
			if item != "4" {
				t.Error("map error on second", item)
			}
		case 2:
			count = count + 1
			if item != "9" {
				t.Error("map error on third", item)
			}
		default:
			t.Error("map error on index", index)
		}
	}

	if count != 3 {
		t.Error("wrong count when map", count)
	}
}
