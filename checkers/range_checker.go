package checkers

import (
	"fmt"
	"math"
	"strings"
)

type RangeChecker[T RealNum] struct {
	Left          *float64
	Right         *float64
	LeftIncluded  bool
	RightIncluded bool
}

func (rc *RangeChecker[T]) CheckInt(val int64, bitSize int8) (bool, string) {
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

	return isLeftOk && isRightOk, rc.Description(val)
}

func (rc *RangeChecker[T]) CheckUint(val uint64, bitSize int8) (bool, string) {
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

	return isLeftOk && isRightOk, rc.Description(val)
}

func (rc *RangeChecker[T]) CheckFloat(val float64, bitSize int8) (bool, string) {
	if math.IsNaN(val) {
		return false, rc.Description(val)
	}

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

	return isLeftOk && isRightOk, rc.Description(val)
}

func (rc *RangeChecker[T]) Description(val any) string {
	return fmt.Sprintf("Value %v is not in the interval %s", val, rc.RangeString())
}

func (rc *RangeChecker[T]) RangeString() string {
	if rc.Left == nil && rc.Right == nil {
		return "(-Inf, +Inf)"
	}

	sb := strings.Builder{}

	if rc.LeftIncluded && rc.Left != nil {
		fmt.Fprintf(&sb, "[")
	} else {
		fmt.Fprintf(&sb, "(")
	}

	if rc.Left != nil {
		fmt.Fprintf(&sb, "%v, ", *rc.Left)
	} else {
		fmt.Fprint(&sb, "-Inf, ")
	}

	if rc.Right != nil {
		fmt.Fprintf(&sb, "%v", *rc.Right)
	} else {
		fmt.Fprint(&sb, "+Inf")
	}

	if rc.RightIncluded && rc.Right != nil {
		fmt.Fprintf(&sb, "]")
	} else {
		fmt.Fprintf(&sb, ")")
	}

	return sb.String()
}
