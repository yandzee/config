package checkers

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type RealNum interface {
	Signed | Unsigned | Float
}

type RealNumChecker interface {
	CheckInt(val int64, bitSize int8) bool
	CheckUint(val uint64, bitSize int8) bool
	CheckFloat(val float64, bitSize int8) bool
}
