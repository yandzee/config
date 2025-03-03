package checkers

import (
	"fmt"

	"github.com/yandzee/config/check"
)

var (
	StrNotEmpty = Fn(func(str string) (bool, string) {
		if len(str) == 0 {
			return false, "String is empty"
		}

		return true, ""
	})
)

func IsPositive[T RealNum]() check.Checker[T] {
	v := float64(0)

	return &RealNumCheckerWrapper[T]{
		Underlying: &RangeChecker[T]{
			Left:         &v,
			LeftIncluded: false,
		},
	}
}

func IsNegative[T RealNum]() check.Checker[T] {
	v := float64(0)

	return &RealNumCheckerWrapper[T]{
		Underlying: &RangeChecker[T]{
			Right:         &v,
			RightIncluded: false,
		},
	}
}

func StrLength(l int) check.Checker[string] {
	return Fn(func(str string) (bool, string) {
		if n := len(str); n != l {
			return false, fmt.Sprintf(
				"String `%s` has wrong length, expected %d, got: %d",
				str,
				l,
				n,
			)
		}

		return true, ""
	})
}
