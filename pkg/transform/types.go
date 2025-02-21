package transform

type Transformer[F any, T any] interface {
	Transform(F) (T, error)
}

// type Pipeline struct {
// 	// transformers Transformer
// }

type Piped2[A Transformer[F, T], B Transformer[T, N], F any, T any, N any] struct {
	First  A
	Second B
}

func Pipe2[F any, T any, N any](
	first Transformer[F, T],
	second Transformer[T, N],
) Transformer[F, N] {
	return &Piped2[Transformer[F, T], Transformer[T, N], F, T, N]{
		First:  first,
		Second: second,
	}
}

func (pt *Piped2[A, B, F, T, N]) Transform(from F) (N, error) {
	var n N

	first, err := pt.First.Transform(from)
	if err != nil {
		return n, err
	}

	return pt.Second.Transform(first)
}

type Piped5[
	A Transformer[_A, _B],
	B Transformer[_B, _C],
	C Transformer[_C, _D],
	D Transformer[_D, _E],
	E Transformer[_E, _F],
	_A, _B, _C, _D, _E, _F any,
] struct {
	TransformerA A
	TransformerB B
	TransformerC C
	TransformerD D
	TransformerE E
}
