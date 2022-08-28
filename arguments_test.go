package lib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectNil(t *testing.T) {
	var v interface{}
	fmt.Println(reflect.TypeOf(v) == nil) // true

	val := detectNil(v)
	assert.Equal(t, true, val)
}

func TestNotNil(t *testing.T) {
	var randomNil interface{}
	s := Arguments{}

	val := s.NotNil(1, "not nil")

	assert.Panics(t, func() { s.NotNil(randomNil, "nil") }, "should panic because empty interface{}")
	assert.Equal(t, 1, val)
}

func TestNotEmpty(t *testing.T) {
	s := Arguments{}
	val := s.NotEmpty("a", "not empty")

	assert.Equal(t, "a", val)
	assert.Panics(t, func() { s.NotEmpty("", "empty") }, "should panic")
}

func TestIsGreaterThan(t *testing.T) {
	a := Arguments{}
	v := a.IsGreaterThan(1, 0, "greater than")

	assert.Equal(t, 1, v)
	assert.Panics(t, func() { a.IsGreaterThan(1, 1, "greater than") }, "should panic")
}

func TestIsLessThan(t *testing.T) {
	a := Arguments{}
	v := a.IsLessThan(1, 2, "less than")

	assert.Panics(t, func() { a.IsLessThan(1, 1, "less than") }, "should panic")
	assert.Equal(t, 1, v)
}

func TestIsGreaterThanOrEqualTo(t *testing.T) {
	a := Arguments{}
	v := a.IsGreaterThanOrEqualTo(1, 0, "greater than or equal to")
	val := a.IsGreaterThanOrEqualTo(1, 1, "greater than or equal to")

	assert.Equal(t, 1, v)
	assert.Equal(t, 1, val)
	assert.Panics(t, func() { a.IsGreaterThanOrEqualTo(1, 2, "greater than or equal to") }, "should panic")
}

func TestIsLessThanOrEqualTo(t *testing.T) {
	a := Arguments{}
	lessThan := a.IsLessThanOrEqualTo(1, 2, "less than or equal to")
	equalTo := a.IsLessThanOrEqualTo(1, 1, "less than or equal to")

	assert.Equal(t, 1, lessThan)
	assert.Equal(t, 1, equalTo)
	assert.Panics(t, func() { a.IsLessThanOrEqualTo(1, 0, "less than or equal to") }, "should panic")
}
