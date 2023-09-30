package ppz_sli

import (
	"fmt"
	"testing"
)

func Test_map(t *testing.T) {
	raw := []int{1, 2, 3}
	power2 := Map(raw, func(el int) string {
		return fmt.Sprintf("%d", el*el)
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
