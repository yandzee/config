package checkers

import (
	"errors"
	"fmt"
	"os"

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

	FileExists = Fn(func(r *result.Result[string]) (bool, string) {
		if _, err := os.Stat(r.Value); errors.Is(err, os.ErrNotExist) {
			return false, fmt.Sprintf("File `%s` does not exist", r.Value)
		}

		return true, ""
	})

	FilesExist = Fn(func(r *result.Result[[]string]) (bool, string) {
		for idx, filePath := range r.Value {
			if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
				return false, fmt.Sprintf("File `%s` (index: %d) does not exist", filePath, idx)
			}
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

func IsBetween[T RealNum, V RealNum](left, right V) check.Checker[T] {
	l := float64(left)
	r := float64(right)

	return &RealNumCheckerWrapper[T]{
		Underlying: &RangeChecker[T]{
			Left:          &l,
			Right:         &r,
			RightIncluded: true,
			LeftIncluded:  true,
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
