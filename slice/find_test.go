package ppz_sli

import "testing"

func Test_find_item(t *testing.T) {
	raw := []int{1, 3, 4, 5, 8, 10}
	count := 0
	gt5, index, ok := Arm(raw).Find(func(el int) bool {
		count = count + 1
		return el > 5
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
	six, index, ok := Arm(raw).Find(func(el int) bool {
		count += 1
		return el == 6
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
