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
	val := s.NotEmpty("a3eef915-78cf-4958-b84f-f0d6f81fe4ad", "not empty")

	assert.Equal(t, "a3eef915-78cf-4958-b84f-f0d6f81fe4ad", val)
	assert.Panics(t, func() { s.NotEmpty("", "empty") }, "should panic")
}

func TestNotWhiteSpace(t *testing.T) {
	s := Arguments{}
	val := s.NotWhiteSpace("a", "not empty")

	assert.Equal(t, "a", val)
	assert.Panics(t, func() { s.NotWhiteSpace("", "empty") }, "should panic")
}

func TestIsGreaterThan(t *testing.T) {
	a := Arguments{}
	v := a.IsGreaterThan(1, 0, "greater than")

	assert.Equal(t, 1, v)
	assert.Panics(t, func() { a.IsGreaterThan(1, 2, "greater than") }, "should panic")
}

func TestIsLessThan(t *testing.T) {
	a := Arguments{}
	v := a.IsLessThan(1, 2, "less than")

	assert.Panics(t, func() { a.IsLessThan(2, 1, "less than") }, "should panic")
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

func TestIsBetween(t *testing.T) {
	a := Arguments{}

	fromInclusive := a.IsBetween(1, 1, 2, "")
	toInclusive := a.IsBetween(2, 1, 2, "")
	between := a.IsBetween(2, 1, 3, "")

	assert.Equal(t, 1, fromInclusive)
	assert.Equal(t, 2, toInclusive)
	assert.Equal(t, 2, between)
	assert.Panics(t, func() { a.IsBetween(5, 1, 2, "should fail") }, "should fail")
	assert.Panics(t, func() { a.IsBetween(1, 2, 3, "should fail") }, "should fail")
}

const rawBase64 = `iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAAAXNSR0IArs4c6QAAA2RJREFUaEPtmr9r8kAYx5+CKNKhoINQoYMUOri5qnM7uzjYURBsK3RVCtriLEi1WHAqLXToH+BeNxEchdKpFgctOIh0ystzL/G9pEnucrm0pryZPC957vt5fp3E2xoMBkooFIJAIABevD4/P+Hj4wO2Xl9fFfywv78POzs7nmJZLBbw8vICGIityWSibG9vky+8BKNCoOblcvkXZHd3F+iJTY+MXuv7+/s/EMwpL8AYafwCsukwZo42BNlUGKtsMQXZNBhWyluCbAoMCwJ1MkF+GoYHghvkp2B4IWyBfDeMHQjbIN8FYxdCCMRtGBEIYRC3YEQhHIHIhnEC4RhEFoxTCCkgTmFkQEgDEYWRBSEVxC6MTAjpILwwsiFcAWHBuAHhGogZjFsQroLoYXDs5ssNrp/xKEL0UqOAz7v5huY/CE+E6JrwbGoZFbbnit1KsFsw0muERyjPPTypS98jFcSOQDv38kBJAxERJvKMGZQUECeCnDwrNbVkCJFhw1FEZAhQverUljCI04WNct2JTSEQJwuyOpCobdsgoguxAOh5kTVsgdAL+P1+OD8/h06ns9ZweXkJFxcXZDwejyGbzcJoNIJCoQCNRgOCwaAlz2q1+mKzXC5DvV5n2uQG0XtpPp/D6ekpVKtVODg40AhUBaXTachkMkQcfs7lcpYgeptGjjOzyQViFGr0OEJcX19DOBzWCMS5UqkEzWaTQPb7fbi7u2NGxcimuraiKFCpVExtMkHM8tVKHM5dXV3B/f09gaTH7XabjOm5k5MTeHx8hNlsZgiMGp6enuDh4YHcp7eJY+G/3lDI8fHxOhKHh4cacXQE9J5GSLyKxSJJN6yrZDJJnjez2ev14Pb2FtARkUiE1CCdEcJ/hqKYyWSyThd6PBwONZ7VL4q1gAAojm4QLJvdbhfy+TzE43GYTqdsEJH2R9cFpohZaqn1hN7HqKkpZtQFjGze3NyQcye4BnZC9XlpBwZor+MidNj19aQKRPGpVGrdsvUwZjZ9Ph+pmefnZ2i1WqSta0B4I0G3V0wRdRyNRokoq/ZLzx0dHa1rJJFIaNo0yyY2iFgsBmdnZ+Qg0BrE7qEa/eal3/SMNkT0OO4peKkbJEZL7Vp7e3uaDZFls1arwdvbG3nNRA7V/JpjTr/l4NkfxDvpqyH6vqsAAAAASUVORK5CYII=`
const invalidBase64 = `iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAAAXNSR0IArs4c6QAAA2RJREFUaEPtmr9r8kAYx5+CKNKhoINQoYMUOri5qnM7uzjYURBsK3RVCtriLEi1WHAqLXToH+BeNxEchdKpFgctOIh0ystzL/G9pEnucrm0pryZPC957vt5fp3E2xoMBkooFIJAIABevD4/P+Hj4wO2Xl9fFfywv78POzs7nmJZLBbw8vICGIityWSibG9vky+8BKNCoOblcvkXZHd3F+iJTY+MXuv7+/s/EMwpL8AYafwCsukwZo42BNlUGKtsMQXZNBhWyluCbAoMCwJ1MkF+GoYHghvkp2B4IWyBfDeMHQjbIN8FYxdCCMRtGBEIYRC3YEQhHIHIhnEC4RhEFoxTCCkgTmFkQEgDEYWRBSEVxC6MTAjpILwwsiFcAWHBuAHhGogZjFsQroLoYXDs5ssNrp/xKEL0UqOAz7v5huY/CE+E6JrwbGoZFbbnit1KsFsw0muERyjPPTypS98jFcSOQDv38kBJAxERJvKMGZQUECeCnDwrNbVkCJFhw1FEZAhQverUljCI04WNct2JTSEQJwuyOpCobdsgoguxAOh5kTVsgdAL+P1+OD8/h06ns9ZweXkJFxcXZDwejyGbzcJoNIJCoQCNRgOCwaAlz2q1+mKzXC5DvV5n2uQG0XtpPp/D6ekpVKtVODg40AhUBaXTachkMkQcfs7lcpYgeptGjjOzyQViFGr0OEJcX19DOBzWCMS5UqkEzWaTQPb7fbi7u2NGxcimuraiKFCpVExtMkHM8tVKHM5dXV3B/f09gaTH7XabjOm5k5MTeHx8hNlsZgiMGp6enuDh4YHcp7eJY+G/3lDI8fHxOhKHh4cacXQE9J5GSLyKxSJJN6yrZDJJnjez2ev14Pb2FtARkUiE1CCdEcJ/hqKYyWSyThd6PBwONZ7VL4q1gAAojm4QLJvdbhfy+TzE43GYTqdsEJH2R9cFpohZaqn1hN7HqKkpZtQFjGze3NyQcye4BnZC9XlpBwZor`

func TestIsValidBase64(t *testing.T) {
	a := Arguments{}

	rawBase := a.IsValidBase64(rawBase64, "it works")

	assert.Equal(t, rawBase64, rawBase)
	assert.Panics(t, func() { a.IsValidBase64(invalidBase64, "fails") }, "fails")
}
