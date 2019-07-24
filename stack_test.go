package adt_test

import (
	"testing"

	"github.com/braddle/adt"
	"github.com/stretchr/testify/suite"
)

type StackSuite struct {
	suite.Suite
	stack *adt.Stack
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

func (s *StackSuite) SetupTest() {
	s.stack = adt.NewStack()
}

func (s *StackSuite) TestEmpty() {
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestNotEmpty() {
	s.stack.Push("red")

	s.False(s.stack.IsEmpty())
}

func (s *StackSuite) TestSize() {
	s.Equal(0, s.stack.Size())

	s.stack.Push("red")
	s.Equal(1, s.stack.Size())

	s.stack.Push("Blue")
	s.Equal(2, s.stack.Size())
}

func (s *StackSuite) TestPop() {
	s.stack.Push("Red")
	s.stack.Push("Blue")

	s.Equal("Blue", s.stack.Pop())
	s.Equal(1, s.stack.Size())

	s.Equal("Red", s.stack.Pop())
	s.True(s.stack.IsEmpty())
}
