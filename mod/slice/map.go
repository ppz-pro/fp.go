package ppz_sli

func Map3[Input any, Output any](inputs []Input, cb func(Input, int, []Input) Output) Slice[Output] {
	result := make_container[Output](len(inputs))
	for index, item := range inputs {
		result[index] = cb(item, index, inputs)
	}
	return result
}

func Map2[Input any, Output any](inputs []Input, cb func(Input, int) Output) Slice[Output] {
	return Map3(inputs, func(el Input, index int, _ []Input) Output {
		return cb(el, index)
	})
}

func Map[Input any, Output any](inputs []Input, cb func(Input) Output) Slice[Output] {
	return Map3(inputs, func(el Input, _ int, _ []Input) Output {
		return cb(el)
	})
}
