package source

import (
	"os"
)

type EnvVarSource struct {
	VarName string
}

func (evs *EnvVarSource) Lookup() (string, bool, error) {
	v, ok := os.LookupEnv(evs.VarName)
	return v, ok, nil
}

func (evs *EnvVarSource) Name() string {
	return evs.VarName
}

func (evs *EnvVarSource) Kind() string {
	return "env"
}
