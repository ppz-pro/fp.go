package ppz_sli

import (
	"fmt"
	"testing"
)

func Test_reduce(t *testing.T) {
	result := "666-"
	acc := Reduce[int, string](
		&result,
		[]int{1, 2, 3, 10, 20, 30},
		func(acc *string, el int) {
			*acc += fmt.Sprintf("%d", el)
		},
	)
	if acc != "666-123102030" {
		t.Error("slice.Reduce error", acc)
	}
}
