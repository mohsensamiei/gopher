package closer

func Add(fn Close) {
	functions = append(functions, fn)
}
