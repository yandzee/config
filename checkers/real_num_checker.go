package checkers

import "fmt"

type RealNumCheckerWrapper[T RealNum] struct {
	Underlying RealNumChecker
}

func (w *RealNumCheckerWrapper[T]) Check(v T) (bool, string) {
	switch val := any(v).(type) {
	case int:
		return w.Underlying.CheckInt(int64(val), 0)
	case int8:
		return w.Underlying.CheckInt(int64(val), 8)
	case int16:
		return w.Underlying.CheckInt(int64(val), 16)
	case int32:
		return w.Underlying.CheckInt(int64(val), 32)
	case int64:
		return w.Underlying.CheckInt(val, 64)
	case uint:
		return w.Underlying.CheckUint(uint64(val), 0)
	case uint8:
		return w.Underlying.CheckUint(uint64(val), 8)
	case uint16:
		return w.Underlying.CheckUint(uint64(val), 16)
	case uint32:
		return w.Underlying.CheckUint(uint64(val), 32)
	case uint64:
		return w.Underlying.CheckUint(val, 64)
	case float32:
		return w.Underlying.CheckFloat(float64(val), 32)
	case float64:
		return w.Underlying.CheckFloat(val, 64)
	}

	return false, fmt.Sprintf("RealNumCheckerWrapper: unknown value %v (%T)", v, v)
}
