package configurator

import "log/slog"

type DescriptorFlag uint8

const (
	DescFlagRequired DescriptorFlag = 1 << iota
	DescFlagPresented
	DescFlagNotPresented
	DescFlagDefaulted
	DescFlagParseError
	DescFlagCustomError
	DescFlagLookupError
)

func (df DescriptorFlag) IsRequired() bool {
	return df&DescFlagRequired != 0
}

func (df DescriptorFlag) IsPresented() bool {
	return df&DescFlagPresented != 0
}

func (df DescriptorFlag) IsDefaulted() bool {
	return df&DescFlagDefaulted != 0
}

func (df DescriptorFlag) IsNotPresented() bool {
	return df&DescFlagNotPresented != 0
}

func (df DescriptorFlag) IsParseError() bool {
	return df&DescFlagParseError != 0
}

func (df DescriptorFlag) IsCustomError() bool {
	return df&DescFlagCustomError != 0
}

func (df DescriptorFlag) IsLookupError() bool {
	return df&DescFlagLookupError != 0
}

func (df *DescriptorFlag) Add(flags ...DescriptorFlag) {
	for _, f := range flags {
		*df |= f
	}
}

func (df *DescriptorFlag) Remove(flags ...DescriptorFlag) {
	for _, f := range flags {
		*df &= ^f
	}
}

func (df DescriptorFlag) LogAttrs() []any {
	attrs := []any{}

	pairs := []struct {
		string
		bool
	}{
		{"is-required", df.IsRequired()},
		{"is-presented", df.IsPresented()},
		{"is-not-presented", df.IsNotPresented()},
		{"is-defaulted", df.IsDefaulted()},
		{"is-parse-error", df.IsParseError()},
		{"is-custom-error", df.IsCustomError()},
		{"is-lookup-error", df.IsLookupError()},
	}

	for _, pair := range pairs {
		if !pair.bool {
			continue
		}

		attrs = append(attrs, slog.Bool(pair.string, pair.bool))
	}

	return attrs
}
