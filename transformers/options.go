package transformers

func CoerceOptions[T any](opts []any) []T {
	casted := make([]T, 0, len(opts))

	for _, opt := range opts {
		castedOpt, ok := opt.(T)
		if !ok {
			continue
		}

		casted = append(casted, castedOpt)
	}

	return casted
}
