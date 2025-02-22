package source

type StrSource struct {
	Str       string
	Presented bool
}

func (ss *StrSource) Lookup() (string, bool, error) {
	return ss.Str, ss.Presented, nil
}

func (ss *StrSource) Name() string {
	return ss.Str
}

func (ss *StrSource) Kind() string {
	return "str"
}
