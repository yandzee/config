package result

import (
	"log/slog"
	"strconv"
)

type ResultFlag uint8

const (
	FlagRequired ResultFlag = 1 << iota
	FlagPresented
	FlagDefaulted
	FlagCheckFailed
	FlagTransformError
	FlagDefaulterError
	FlagLookupError
)

type FlagPair struct {
	Name    string
	Enabled bool
}

func (rf ResultFlag) IsRequired() bool {
	return rf&FlagRequired != 0
}

func (rf ResultFlag) IsPresented() bool {
	return rf&FlagPresented != 0
}

func (rf ResultFlag) IsDefaulted() bool {
	return rf&FlagDefaulted != 0
}

func (rf ResultFlag) IsCheckFailed() bool {
	return rf&FlagCheckFailed != 0
}

func (rf ResultFlag) IsTransformError() bool {
	return rf&FlagTransformError != 0
}

func (rf ResultFlag) IsDefaulterError() bool {
	return rf&FlagDefaulterError != 0
}

func (rf ResultFlag) IsLookupError() bool {
	return rf&FlagLookupError != 0
}

func (rf ResultFlag) String() string {
	return strconv.FormatInt(int64(rf), 2)
}

func (rf *ResultFlag) Add(flags ...ResultFlag) {
	for _, f := range flags {
		*rf |= f
	}
}

func (rf *ResultFlag) Remove(flags ...ResultFlag) {
	for _, f := range flags {
		*rf &= ^f
	}
}

func (rf ResultFlag) Pairs(all ...bool) []FlagPair {
	pairs := []FlagPair{
		{"is-required", rf.IsRequired()},
		{"is-presented", rf.IsPresented()},
		{"is-defaulted", rf.IsDefaulted()},
		{"is-check-failed", rf.IsCheckFailed()},
		{"is-transform-error", rf.IsTransformError()},
		{"is-defaulter-error", rf.IsDefaulterError()},
		{"is-lookup-error", rf.IsLookupError()},
	}

	if len(all) > 0 && all[0] {
		return pairs
	}

	filtered := []FlagPair{}

	for _, pair := range pairs {
		if !pair.Enabled {
			continue
		}

		filtered = append(filtered, pair)
	}

	return filtered

}

func (rf ResultFlag) LogAttrs() []any {
	attrs := []any{}

	for _, flag := range rf.Pairs() {
		if !flag.Enabled {
			continue
		}

		attrs = append(attrs, slog.Bool(flag.Name, flag.Enabled))
	}

	return attrs
}
