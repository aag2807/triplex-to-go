package lib

import "reflect"

type State struct{}

// IsTrue panics if stateQuery is false.
func (s *State) IsTrue(stateQuery bool, errormsg string) bool {
	if !stateQuery {
		panic(errormsg)
	}
	return true
}

// IsFalse panics if stateQuery is true.
func (s *State) IsFalse(stateQuery bool, errormsg string) bool {
	if stateQuery {
		panic(errormsg)
	}

	return true
}

// Equals verifies deep equal between both arguments, will panic if they dont match.
func (s *State) Equals(arg1, arg2 interface{}, errormsg string) bool {
	if arg1 != arg2 || !reflect.DeepEqual(arg1, arg2) {
		panic(errormsg)
	}
	return true
}
