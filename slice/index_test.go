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
	return len(a) == len(b) && cap(a) == cap(b)
}

func Test_each(t *testing.T) {
	raw := []int{1, 2, 3}

	test := make_container[int](len(raw))
	count := 0

	Arm(raw).Each(func(item Item[int]) {
		if item.Index != count {
			t.Error("index error", item.Index, count)
		}
		test[item.Index] = item.El
		count = count + 1
	})

	if !is_same_slice([]int{1, 2, 3}, test) {
		t.Error("wrong el after each", test)
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

func Test_find_item(t *testing.T) {
	raw := []int{1, 3, 4, 5, 8, 10}
	count := 0
	gt5, index, ok := Arm(raw).Find(func(item Item[int]) bool {
		count = count + 1
		return item.El > 5
	})

	if gt5 != 8 {
		t.Error("gt5 isn't 8", gt5)
	}
	if index != 4 {
		t.Error("gt5 isn't on forth", index)
	}
	if !ok {
		t.Error("no gt5", ok)
	}
	if count != 5 {
		t.Error("wrong count on found", count)
	}

	count = 0
	six, index, ok := Arm(raw).Find(func(item Item[int]) bool {
		count += 1
		return item.El == 6
	})
	if six != 0 {
		t.Error("found six", six)
	}
	if index != 0 {
		t.Error("fount six index", index)
	}
	if ok {
		t.Error("found six", ok)
	}
	if count != 6 {
		t.Error("wrong count on not found", count)
	}
}

func Test_reverse(t *testing.T) {
	raw := []int{1, 2, 3, 4, 6}
	reversed := Arm(raw).Reverse()
	if !is_same_slice(reversed, []int{6, 4, 3, 2, 1}) {
		t.Error("reverse error", reversed)
	}
}

func Test_every(t *testing.T) {
	raw := []int{2, 4, 6, 8, 10}
	count := 0
	is_all_even := Arm(raw).Every(func(item Item[int]) bool {
		count += 1
		return item.El%2 == 0
	})
	if !is_all_even {
		t.Error("slice.every not all even")
	}
	if count != 5 {
		t.Error("slice.every iteration time is not 5", count)
	}
}

func Test_some(t *testing.T) {
	raw := []int{1, 3, 6, 2, 1, 10, 11}
	count := 0

	has6 := Arm(raw).Some(func(item Item[int]) bool {
		count += 1
		return item.El == 6
	})

	if !has6 {
		t.Error("slice.Some no 6")
	}
	if count != 3 {
		t.Error("slice.Some times of test is not 3", count)
	}
}

func Test_reduce(t *testing.T) {
	acc := Reduce[int, string](
		"666-",
		[]int{1, 2, 3, 10, 20, 30},
		func(acc string, item Item[int]) string {
			return acc + fmt.Sprintf("%d", item.El)
		},
	)
	if acc != "666-123102030" {
		t.Error("slice.Reduce error", acc)
	}
}
