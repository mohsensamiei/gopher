package closer

func Defer() {
	for _, fn := range functions {
		func(fn Close) {
			defer func() {
				_ = recover()
			}()
			_ = fn()
		}(fn)
	}
}
