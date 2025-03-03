package check

import "github.com/yandzee/config/result"

type Checker[T any] interface {
	Check(r *result.Result[T]) (bool, string)
}

func Run[T any](chkrs ...Checker[T]) bool {
	return true
}
