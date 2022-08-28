package lib

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestIsTrue(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	s := State{}

	g.Expect(s.IsTrue(1 == 1, "true because 1 = 1")).To(gomega.BeTrue(), "should not panic")
	assert.Panics(t, func() { s.IsTrue(1 == 2, "false because 1 != 2") }, "Should panic")
}

func TestIsFalse(t *testing.T) {
	s := State{}
	val := s.IsFalse(1 == 2, "false because 1 != 2")

	assert.Equal(t, true, val)
	assert.Panics(t, func() { s.IsFalse(1 == 1, "true because 1 = 1") }, "Should panic")
}

func TestIsEqual(t *testing.T) {
	s := State{}
	val := s.Equals(1, 1, "true because 1 = 1")

	assert.Equal(t, true, val)
	assert.Panics(t, func() { s.Equals(1, 2, "false because 1 != 2") }, "Should panic")
}
