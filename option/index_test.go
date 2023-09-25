package option

import "testing"

func Test_option_with_value(t *testing.T) {
	option := New[int](1)

	has_executed := false

	option.Match(
		func(v int) {
			has_executed = true
			if v != 1 {
				t.Error("option value error ", v)
			}
		},
		func() {
			t.Error("no_value method should not be executed on option without value")
		},
	)

	if !has_executed {
		t.Error("has_value not executed on option with value")
	}
}

func Test_option_without_value(t *testing.T) {
	no_value := Nil[int]()

	has_executed := false

	no_value.Match(
		func(_ int) {
			t.Error("has_value shouldn't be executed on option without value")
		},
		func() {
			has_executed = true
		},
	)

	if !has_executed {
		t.Error("no_value not executed no option without value")
	}
}
