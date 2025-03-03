package checkers

import (
	"github.com/yandzee/config/check"
)

type RangeChecker[T RealNum] struct {
	Left          *float64
	Right         *float64
	LeftIncluded  bool
	RightIncluded bool
}

func (rc *RangeChecker[T]) CheckInt(val int64, bitSize int8) bool {
	isLeftOk, isRightOk := true, true

	if rc.Left != nil {
		left := int64(*rc.Left)

		if rc.LeftIncluded {
			isLeftOk = val >= left
		} else {
			isLeftOk = val > left
		}
	}

	if rc.Right != nil {
		right := int64(*rc.Right)

		if rc.RightIncluded {
			isRightOk = val <= right
		} else {
			isRightOk = val < right
		}
	}

	return isLeftOk && isRightOk
}

func (rc *RangeChecker[T]) CheckUint(val uint64, bitSize int8) bool {
	isLeftOk, isRightOk := true, true

	if rc.Left != nil {
		left := uint64(*rc.Left)

		if rc.LeftIncluded {
			isLeftOk = val >= left
		} else {
			isLeftOk = val > left
		}
	}

	if rc.Right != nil {
		right := uint64(*rc.Right)

		if rc.RightIncluded {
			isRightOk = val <= right
		} else {
			isRightOk = val < right
		}
	}

	return isLeftOk && isRightOk
}

func (rc *RangeChecker[T]) CheckFloat(val float64, bitSize int8) bool {
	isLeftOk, isRightOk := true, true

	if rc.Left != nil {
		left := *rc.Left

		if rc.LeftIncluded {
			isLeftOk = val >= left
		} else {
			isLeftOk = val > left
		}
	}

	if rc.Right != nil {
		right := *rc.Right

		if rc.RightIncluded {
			isRightOk = val <= right
		} else {
			isRightOk = val < right
		}
	}

	return isLeftOk && isRightOk
}

func IsPositive[T RealNum]() check.Checker[T] {
	left := float64(0)

	return &RealNumCheckerWrapper[T]{
		Underlying: &RangeChecker[T]{
			Left:         &left,
			LeftIncluded: false,
		},
	}
}
