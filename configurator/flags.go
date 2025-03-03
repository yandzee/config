package configurator

import (
	"log/slog"
	"strconv"
)

type DescriptorFlag uint8

const (
	DescFlagRequired DescriptorFlag = 1 << iota
	DescFlagPresented
	DescFlagDefaulted
	DescFlagCheckFailed
	DescFlagTransformError
	DescFlagCustomError
	DescFlagLookupError
)

type DescriptorFlagPair struct {
	Name    string
	Enabled bool
}

func (df DescriptorFlag) IsRequired() bool {
	return df&DescFlagRequired != 0
}

func (df DescriptorFlag) IsPresented() bool {
	return df&DescFlagPresented != 0
}

func (df DescriptorFlag) IsDefaulted() bool {
	return df&DescFlagDefaulted != 0
}

func (df DescriptorFlag) IsCheckFailed() bool {
	return df&DescFlagCheckFailed != 0
}

func (df DescriptorFlag) IsTransformError() bool {
	return df&DescFlagTransformError != 0
}

func (df DescriptorFlag) IsCustomError() bool {
	return df&DescFlagCustomError != 0
}

func (df DescriptorFlag) IsLookupError() bool {
	return df&DescFlagLookupError != 0
}

func (df DescriptorFlag) String() string {
	return strconv.FormatInt(int64(df), 2)
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

func (df DescriptorFlag) Pairs(all ...bool) []DescriptorFlagPair {
	pairs := []DescriptorFlagPair{
		{"is-required", df.IsRequired()},
		{"is-presented", df.IsPresented()},
		{"is-defaulted", df.IsDefaulted()},
		{"is-check-failed", df.IsCheckFailed()},
		{"is-transform-error", df.IsTransformError()},
		{"is-custom-error", df.IsCustomError()},
		{"is-lookup-error", df.IsLookupError()},
	}

	if len(all) > 0 && all[0] {
		return pairs
	}

	filtered := []DescriptorFlagPair{}

	for _, pair := range pairs {
		if !pair.Enabled {
			continue
		}

		filtered = append(filtered, pair)
	}

	return filtered

}

func (df DescriptorFlag) LogAttrs() []any {
	attrs := []any{}

	for _, flag := range df.Pairs() {
		if !flag.Enabled {
			continue
		}

		attrs = append(attrs, slog.Bool(flag.Name, flag.Enabled))
	}

	return attrs
}
