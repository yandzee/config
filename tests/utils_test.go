package tests

type StrSource struct {
	Str       string
	Presented bool
	Error     error
}

func NewStr(str string, err error, presented ...bool) *StrSource {
	return &StrSource{
		Str:       str,
		Presented: len(presented) == 0 || presented[0],
		Error:     err,
	}
}

func (ss *StrSource) Lookup() (string, bool, error) {
	return ss.Str, ss.Presented, ss.Error
}

func (ss *StrSource) Name() string {
	return ss.Str
}

func (ss *StrSource) Kind() string {
	return "str"
}
