package environment

import (
	"os"
	"strings"
)

type Environment string

func Current() Environment {
	osEnv := strings.ToLower(os.Getenv("ENV"))
	for _, env := range All() {
		if string(env) == osEnv {
			return env
		}
	}

	return Default
}

func (e Environment) HasParent() bool {
	return e.Parent() != e
}

func All() []Environment {
	return []Environment{Production, Staging, Development, Default}
}

func (e Environment) Parent() Environment {
	return Default
}
