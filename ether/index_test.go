package ether

import "testing"

func Test_left(t *testing.T) {
	left_ether := New_left[int, string](1)

	has_executed := false

	left_ether.Match(
		func(left_value int) {
			has_executed = true
			if left_value != 1 {
				t.Error("wrong left value", left_value)
			}
		},
		func(right_value string) {
			t.Error("right shouldn't be executed no Ether with left value")
		},
	)

	if !has_executed {
		t.Error("left not executed on Ether with left value")
	}
}

func Test_right(t *testing.T) {
	right_ether := New_right[int, string]("ppz")

	has_executed := false

	right_ether.Match(
		func(left_value int) {
			t.Error("left shouldn't be executed no Ether with right value")
		},
		func(right_value string) {
			has_executed = true
			if right_value != "ppz" {
				t.Error("wrong right value", right_value)
			}
		},
	)

	if !has_executed {
		t.Error("right not executed on Ether with right value")
	}
}
