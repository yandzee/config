package common

type KVOptions map[string]any

func (kv KVOptions) GetString(key string) string {
	return kv.GetStringOr(key, "")
}

func (kv KVOptions) GetStringSliceOr(key string, def []string) []string {
	if kv == nil {
		return def
	}

	s, ok := kv[key].([]string)
	if ok {
		return s
	}

	return def
}

func (kv KVOptions) GetBoolOr(key string, def bool) bool {
	if kv == nil {
		return def
	}

	b, ok := kv[key].(bool)
	if ok {
		return b
	}

	return def
}

func (kv KVOptions) GetStringOr(key, def string) string {
	if kv == nil {
		return def
	}

	s, ok := kv[key].(string)
	if ok {
		return s
	}

	return def
}

func (kv KVOptions) Set(key string, value any) {
	if kv == nil {
		return
	}

	kv[key] = value
}
