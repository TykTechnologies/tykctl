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

func NewGenericHTTPError(body string) GenericHTTPError {
	return GenericHTTPError{Body: body}
}

type GenericHTTPError struct {
	Body string
}

func (g GenericHTTPError) Error() string {
	return g.Body
}
