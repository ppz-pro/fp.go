package ppz_sli

func Reduce3[El any, Result any](
	result *Result,
	slice []El,
	cb func(*Result, El, int, []El),
) Result {
	for index, el := range slice {
		cb(result, el, index, slice)
	}
	return *result
}

func Reduce2[El any, Result any](
	result *Result,
	slice []El,
	cb func(*Result, El, int),
) Result {
	return Reduce3[El, Result](
		result,
		slice,
		func(result *Result, el El, index int, _ []El) {
			cb(result, el, index)
		},
	)
}

func Reduce[El any, Result any](
	result *Result,
	slice []El,
	cb func(*Result, El),
) Result {
	return Reduce3[El, Result](
		result,
		slice,
		func(result *Result, el El, _ int, _ []El) {
			cb(result, el)
		},
	)
}
