package transformers

import "fmt"

func CoerceOptions[O, T any](opts []O) []T {
	casted := make([]T, 0, len(opts))

	for _, opt := range opts {
		castedOpt, ok := any(opt).(T)
		if !ok {
			panic(
				fmt.Sprintf(
					"CoerceOptions: cannot coerce value `%v` (%T) to type %T",
					opt,
					opt,
					castedOpt,
				),
			)
		}

		casted = append(casted, castedOpt)
	}

	return casted
}
