package source

type StringSource interface {
	Lookup() (string, bool, error)
	Kind() string
	Name() string
}
