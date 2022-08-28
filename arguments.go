package lib

import (
	"reflect"
)

type Arguments struct{}

// detectNil returns true if the value is nil meant for internal use within the package
func detectNil[T any](arg T) bool {
	if reflect.ValueOf(arg).Kind() == reflect.Ptr && reflect.ValueOf(arg).IsNil() || reflect.TypeOf(arg) == nil {
		return true
	}

	return false
}

// NotNil returns the value if the value if not null/nil
func (a *Arguments) NotNil(arg interface{}, errormsg string) interface{} {
	if detectNil(arg) {
		panic(errormsg)
	}
	return arg
}

func (a *Arguments) NotEmpty(arg string, errormsg string) string {
	if arg == "" {
		panic(errormsg)
	}
	return arg
}

func (a *Arguments) IsGreaterThan(arg, comprarer int, errormsg string) int {
	if arg <= comprarer {
		panic(errormsg)
	}
	return arg
}

func (a *Arguments) IsLessThan(arg, comprarer int, errormsg string) int {
	if arg >= comprarer {
		panic(errormsg)
	}
	return arg
}

func (a *Arguments) IsGreaterThanOrEqualTo(arg, comprarer int, errormsg string) int {
	if arg < comprarer {
		panic(errormsg)
	}
	return arg
}

func (a *Arguments) IsLessThanOrEqualTo(arg, comprarer int, errormsg string) int {
	if arg > comprarer {
		panic(errormsg)
	}
	return arg
}
