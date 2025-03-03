package check

type Checker[T any] interface {
	Check(v T) (bool, string)
}

func Run[T any](chkrs ...Checker[T]) bool {
	return true
}
