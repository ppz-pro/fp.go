package fp

import "testing"

func Test_foreach(t *testing.T) {
	raw := []int{1, 2, 3}
	test := make_container[int](raw)
	sum := 0
	Foreach(raw, func(item int) {
		sum = sum + 1
		test = append(test, item)
	})
}
