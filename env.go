package main

import (
	"cmp"
	"os"

	"github.com/konflux-ci/namespace-lister/internal/constants"
)

func getAddress() string {
	return cmp.Or(os.Getenv(constants.EnvAddress), constants.DefaultAddr)
}
