package parse

type FnOptions map[string]any
type Fn[T any] func(string, FnOptions) (T, error)

func (fno FnOptions) GetString(key string) string {
	return fno.GetStringOr(key, "")
}

func (fno FnOptions) GetStringSliceOr(key string, def []string) []string {
	if fno == nil {
		return def
	}

	s, ok := fno[key].([]string)
	if ok {
		return s
	}

	return def
}

func (fno FnOptions) GetBoolOr(key string, def bool) bool {
	if fno == nil {
		return def
	}

	b, ok := fno[key].(bool)
	if ok {
		return b
	}

	return def
}

func (fno FnOptions) GetStringOr(key, def string) string {
	if fno == nil {
		return def
	}

	s, ok := fno[key].(string)
	if ok {
		return s
	}

	return def
}

func (fno FnOptions) Set(key string, value any) {
	if fno == nil {
		return
	}

	fno[key] = value
}
