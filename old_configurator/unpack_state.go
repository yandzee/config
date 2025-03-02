package configurator

import "github.com/yandzee/config/transform"

type UnpackState struct {
	Value          any
	DefaulterError error
	IsInitialized  bool
	IsDefaulted    bool
	IsDefaulterSet bool
}

func (us *UnpackState) GetValue() (any, error) {
	if !us.IsInitialized {
		return us.Value, transform.ErrNoValue
	}

	return us.Value, nil
}

func (us *UnpackState) SetValue(v any) error {
	us.Value = v
	return nil
}
