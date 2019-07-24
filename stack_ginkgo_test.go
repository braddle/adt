package adt_test

import (
	"testing"

	"github.com/braddle/adt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	Describe("Emptiness", func() {
		Context("when the stack does not contain items", func() {
			It("should be empty", func() {
				stack := adt.NewStack()

				Expect(stack.IsEmpty()).To(BeTrue())
			})
		})

		Context("when the stack contains items", func() {
			It("should not be empty", func() {
				stack := adt.NewStack()
				stack.Push("Red")

				Expect(stack.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("Size", func() {
		Context("When items are added to the stack", func() {
			It("should keep track of the size", func() {
				stack := adt.NewStack()

				By("Empty being 0", func() {
					Expect(stack.Size()).To(BeZero())
				})

				By("Adding a first item", func() {
					stack.Push("Red")
					Expect(stack.Size()).To(Equal(1))
				})

				By("Adding a second item", func() {
					stack.Push("Blue")
					Expect(stack.Size()).To(Equal(2))
				})
			})
		})
	})

	Describe("Popping", func() {
		Context("When items are taken off the stack", func() {
			It("Remove items from the top of the stack", func() {

				stack := adt.NewStack()
				stack.Push("Red")
				stack.Push("Blue")

				By("returning the last item added to the stack", func() {
					Expect(stack.Pop()).To(Equal("Blue"))
				})

				By("Decreasing the size", func() {
					Expect(stack.Size()).To(Equal(1))
				})

				By("returning the last item in the stack", func() {
					Expect(stack.Pop()).To(Equal("Red"))
				})

				By("Emptying the stack", func() {
					Expect(stack.IsEmpty()).To(BeTrue())
				})
			})
		})
	})

})

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack")
}
