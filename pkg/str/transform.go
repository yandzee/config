package str

type Transformer[F, T any] interface {
	Transform(F) (T, error)
}

type StringMapper = Transformer[string, string]
type StringSplitter = Transformer[string, []string]
type StringsJoiner = Transformer[[]string, string]
type StringsReducer = Transformer[[]string, []string]

type StringTransformer interface {
	Map(StringMapper) StringTransformer
	Split(StringSplitter) StringsTransformer
	Terminate() StringTerminator
}

type StringTerminator interface {
	Int() (int, error)
	Int8() (int8, error)
	Int16() (int16, error)
	Int32() (int32, error)
	Int64() (int64, error)

	Uint() (uint, error)
	Uint8() (uint8, error)
	Uint16() (uint16, error)
	Uint32() (uint32, error)
	Uint64() (uint64, error)

	Float32() (float32, error)
	Float64() (float64, error)

	Bool() (bool, error)
	Bytes() ([]byte, error)
	String() (string, error)
}

// type StringTerminator2 interface {
// 	Int() Transformer[string, int]
// 	Int8() Transformer[string, int8]
// 	Int16() Transformer[string, int16]
// 	Int32() Transformer[string, int32]
// 	Int64() Transformer[string, int64]
//
// 	Uint() Transformer[string, uint]
// 	Uint8() Transformer[string, uint8]
// 	Uint16() Transformer[string, uint16]
// 	Uint32() Transformer[string, uint32]
// 	Uint64() Transformer[string, uint64]
//
// 	Float32() Transformer[string, float32]
// 	Float64() Transformer[string, float64]
//
// 	Bool() Transformer[string, bool]
// 	Bytes() Transformer[string, []byte]
// 	String() Transformer[string, string]
// }

type StringsTransformer interface {
	Map(StringMapper) StringsTransformer
	Join(StringsJoiner) StringTransformer
	Reduce(StringsReducer) StringsTransformer
	Filter(StringsReducer) StringsTransformer
	Terminate() StringTerminators

	// Execute([]string) ([]string, error)
}

type StringTerminators []StringTerminator

// func Collect(st StringT)

// func (st StringTerminators) Ints() ([]int, error) {
// 	return
// }

func NewTransformer() StringTransformer {
	return nil
}

//
// func Run() {
// }

// type Transform struct{}
//
// func (t *Transform) Map(mapper StrMapper) *Transform {
// 	return "", nil
// }
//
// func (t *Transform) Split(splitter StrSplitter) *Transform {
// 	return nil, nil
// }

//
// str.Transform().
// 	Map(Unbase64).
// 	Map(Unhex)
