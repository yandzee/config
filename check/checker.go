package check

type Checker[T any] interface {
	Check(v T) bool
}

func Run[T any](chkrs ...Checker[T]) bool {
	return true
}
