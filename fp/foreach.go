package fp

func Foreach[Item any](target []Item, cb func(Item)) {
	Foreach_with_index(target, func(item Item, _ int) {
		cb(item)
	})
}

func Foreach_with_index[Item any](target []Item, cb func(Item, int)) {
	Foreach_with_slice(target, func(item Item, index int, _ []Item) {
		cb(item, index)
	})
}

func Foreach_with_slice[Item any](target []Item, cb func(Item, int, []Item)) {
	for index, item := range target {
		cb(item, index, target)
	}
}
