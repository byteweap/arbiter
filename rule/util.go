package rule

// Ternary if condition is true, return trueValue, otherwise return falseValue
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
