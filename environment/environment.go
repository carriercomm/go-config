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

func All() []Environment {
	return []Environment{Production, Staging, Development, Default}
}

func (e Environment) Parent() Environment {
	if e != Default {
		return Default
	}

	return nil
}
