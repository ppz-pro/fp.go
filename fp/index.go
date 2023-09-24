package fp

func make_container[item any](source []item) []item {
	return make([]item, 0, len(source))
}
