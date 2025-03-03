package checkers

import "github.com/yandzee/config/check"

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
