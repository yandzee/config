package check

type Checker[T any] interface {
	Check(v T) bool
}
