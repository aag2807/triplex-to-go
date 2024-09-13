package lib

import (
	"reflect"
	"regexp"
)

type Arguments struct{}

// detectNil returns true if the value is nil meant for internal use within the package
func detectNil[T any](arg T) bool {
	if reflect.ValueOf(arg).Kind() == reflect.Ptr && reflect.ValueOf(arg).IsNil() || reflect.TypeOf(arg) == nil {
		return true
	}

	return false
}

// NotNil returns the value if the value is not null/nil
func (a *Arguments) NotNil(arg interface{}, errormsg string) interface{} {
	if detectNil(arg) {
		panic(errormsg)
	}
	return arg
}

// IsNil makes sure that the argument is nil if not panics
func (a *Arguments) IsNil(arg interface{}, errormsg string) {
	if !detectNil(arg) {
		panic(errormsg)
	}
}

// NotEmpty verifies validity of uuid/Guid and panics if it is false.
func (a *Arguments) NotEmpty(arg string, errormsg string) string {
	uuidRegex, _ := regexp.Compile(`^(\{{0,1}([0-9a-fA-F]){8}-([0-9a-fA-F]){4}-([0-9a-fA-F]){4}-([0-9a-fA-F]){4}-([0-9a-fA-F]){12}\}{0,1})$`)
	if !uuidRegex.MatchString(arg) {
		panic(errormsg)
	}

	return arg
}

// NotWhiteSpace verifies that a string is not empty, if it is empty it will panic.
func (a *Arguments) NotWhiteSpace(arg string, errormsg string) string {
	if arg == "" {
		panic(errormsg)
	}

	return arg
}

// IsGreaterThan panics if the argument is smaller that the comparer non-inclusive.
func (a *Arguments) IsGreaterThan(arg, comparer int, errormsg string) int {
	if arg < comparer {
		panic(errormsg)
	}
	return arg
}

// IsLessThan panics if the argument is bigger that the comparer non-inclusive.
func (a *Arguments) IsLessThan(arg, comparer int, errormsg string) int {
	if arg > comparer {
		panic(errormsg)
	}
	return arg
}

// IsBetween panics if the argument is not in between from and to range
func (a *Arguments) IsBetween(arg, fromInclusive, toInclusive int, errorMessage string) int {
	if arg < fromInclusive || arg > toInclusive {
		panic(errorMessage)
	}

	return arg
}

// IsGreaterThanOrEqualTo panics if the argument is bigger that the comparer inclusive.
func (a *Arguments) IsGreaterThanOrEqualTo(arg, comprarer int, errormsg string) int {
	if arg < comprarer {
		panic(errormsg)
	}
	return arg
}

// IsLessThanOrEqualTo panics if the argument is bigger that the comparer non inclusive.
func (a *Arguments) IsLessThanOrEqualTo(arg, comprarer int, errormsg string) int {
	if arg > comprarer {
		panic(errormsg)
	}
	return arg
}

// isBetween validates that an argument is between a set of numbers inclusively
func (a *Arguments) isBetween(arg, fromInclusive, toInclusive int, errorMsg string) int {
	if arg < fromInclusive {
		panic(errorMsg)
	}

	if arg > toInclusive {
		panic(errorMsg)
	}

	return arg
}

// IsValidBase64 validates if a string  is a valid base 64 it will panic if invalid
func (a *Arguments) IsValidBase64(arg, errorMsg string) string {
	baseRegex, _ := regexp.Compile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)

	if !baseRegex.MatchString(arg) {
		panic(errorMsg)
	}

	return arg
}
