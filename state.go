package lib

import "reflect"

type State struct{}

func (s *State) IsTrue(arg bool, errormsg string) bool {
	if !arg {
		panic(errormsg)
	}
	return true
}

func (s *State) IsFalse(arg bool, errormsg string) bool {
	if arg {
		panic(errormsg)
	}

	return true
}

func (s *State) Equals(arg1, arg2 interface{}, errormsg string) bool {
	if arg1 != arg2 || !reflect.DeepEqual(arg1, arg2) {
		panic(errormsg)
	}
	return true
}
