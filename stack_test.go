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
	// ARRANGE
	s.stack = adt.NewStack()
}

func (s *StackSuite) TestEmpty() {
	// ACT & ASSERT
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestNotEmpty() {
	// ARRANGE
	s.stack.Push("red")

	// ACT
	isEmpty := s.stack.IsEmpty()

	// ASSERT
	s.False(isEmpty)
}

func (s *StackSuite) TestSize() {
	// ACT & ASSERT
	s.Equal(0, s.stack.Size())

	// ARRANGE
	s.stack.Push("red")

	// ACT & ASSERT
	s.Equal(1, s.stack.Size())

	// ARRANGE
	s.stack.Push("Blue")

	// ACT & ASSERT
	s.Equal(2, s.stack.Size())
}

func (s *StackSuite) TestPop() {
	// ARRANGE
	s.stack.Push("Red")
	s.stack.Push("Blue")

	// ACT & ASSERT
	s.Equal("Blue", s.stack.Pop())
	s.Equal(1, s.stack.Size())

	s.Equal("Red", s.stack.Pop())
	s.True(s.stack.IsEmpty())
}
