package internal

import "fmt"

type GenericFlagError struct {
	FlagName string
}

func (g *GenericFlagError) Error() string {
	return fmt.Sprintf("error getting  %s flag", g.FlagName)
}

func NewGenericFlagError(flagName string) GenericFlagError {
	return GenericFlagError{FlagName: flagName}
}
