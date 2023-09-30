package ppz_sli

type Slice[El any] []El

func Arm[El any](el []El) Slice[El] {
	return Slice[El](el)
}

func New[El any](el ...El) Slice[El] {
	return Slice[El](el)
}

func (slice Slice[El]) Disarm() []El {
	return []El(slice)
}

func make_container[item any](size int) []item {
	return make([]item, size)
}
