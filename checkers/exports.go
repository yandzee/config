package checkers

import (
	"fmt"

	"github.com/yandzee/config/check"
	"github.com/yandzee/config/result"
)

var (
	StrNotEmpty = Fn(func(r *result.Result[string]) (bool, string) {
		if len(r.Value) == 0 {
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
	return Fn(func(r *result.Result[string]) (bool, string) {
		if n := len(r.Value); n != l {
			return false, fmt.Sprintf(
				"String `%s` has wrong length, expected %d, got: %d",
				r.Value,
				l,
				n,
			)
		}

		return true, ""
	})
}
