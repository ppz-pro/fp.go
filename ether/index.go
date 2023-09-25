package ether

type Ether[Left any, Right any] struct {
	is_left bool
	left    Left
	right   Right
}

func New_left[Left any, Right any](value Left) Ether[Left, Right] {
	return Ether[Left, Right]{
		is_left: true,
		left:    value,
	}
}
func New_right[Left any, Right any](value Right) Ether[Left, Right] {
	return Ether[Left, Right]{
		is_left: false,
		right:   value,
	}
}

func (ether Ether[Left, Right]) Match(cb_left func(Left), cb_right func(Right)) {
	if ether.is_left {
		cb_left(ether.left)
	} else {
		cb_right(ether.right)
	}
}
