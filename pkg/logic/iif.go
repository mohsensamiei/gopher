package logic

func IIf[T any](logic bool, trueValue func() T, falseValue func() T) T {
	if logic {
		return trueValue()
	} else {
		return falseValue()
	}
}
