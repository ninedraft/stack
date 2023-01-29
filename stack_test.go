package stack_test

import (
	"testing"

	"github.com/ninedraft/stack"
	"golang.org/x/exp/slices"
)

func TestStack_PushPop(t *testing.T) {
	t.Parallel()

	st := stack.Stack[int]{}

	{
		e, ok := st.Pop()
		assertEq(t, false, ok, "poping element (_, ok) = stak.Pop()")
		assertEq(t, 0, e, "poping element (e, _) = stak.Pop()")
	}

	{
		top := st.PopMany(nil, 3)
		assertEq(t, 0, len(top), "poping many elements to a nil distination")
	}

	st.PushMany(1, 2, 3)

	{
		st.Push(4)
		want := stack.Stack[int]{1, 2, 3, 4}
		assertEqSlices(t, want, st, "pushing element")
	}

	{
		e, ok := st.Pop()
		want := 4
		assertEq(t, true, ok, "poping element (_, ok) = stak.Pop()")
		assertEq(t, want, e, "poping element (e, _) = stak.Pop()")
	}

	{
		st.PushMany(4, 5, 6)
		want := stack.Stack[int]{1, 2, 3, 4, 5, 6}
		assertEqSlices(t, want, st, "pushing many elements")
	}

	{
		top := st.PopMany(nil, 3)
		want := []int{6, 5, 4}
		assertEqSlices(t, want, top, "poping many elements to a nil distination")
	}

	{
		dst := []int{10}
		top := st.PopMany(dst, 100)
		want := []int{10, 3, 2, 1}
		assertEqSlices(t, want, top, "poping many elements to a non-nil distination")
	}
}

func TestStack_Peek(t *testing.T) {
	t.Parallel()

	st := stack.Stack[int]{}

	{
		e, ok := st.Peek()
		assertEq(t, false, ok, "peeking element (_, ok) = stak.Peek()")
		assertEq(t, 0, e, "peeking element (e, _) = stak.Peek()")
	}

	{
		top := st.PeekMany(nil, 3)
		assertEq(t, 0, len(top), "peeking many elements to a nil distination")
	}

	st.PushMany(1, 2, 3, 4, 5)

	{
		top, ok := st.Peek()
		want := 5
		assertEq(t, true, ok, "peeking element (_, ok) = stak.Peek()")
		assertEq(t, want, top, "peeking element (e, _) = stak.Peek()")
	}

	{
		top := st.PeekMany(nil, 3)
		want := []int{5, 4, 3}
		assertEqSlices(t, want, top, "peeking many elements to a nil distination")
	}

	{
		dst := []int{10}
		top := st.PeekMany(dst, 100)
		want := []int{10, 5, 4, 3, 2, 1}
		assertEqSlices(t, want, top, "peeking many elements to a non-nil distination")
	}
}

func TestStack_Len(t *testing.T) {
	t.Parallel()

	st := stack.Stack[int]{}
	assertEq(t, 0, st.Len(), "stack{}.Len()")

	st.Push(10)
	st.PushMany(11, 12)

	assertEq(t, 3, st.Len(), "stack{...}.Len()")
}

func TestStack_FreeAfterPop(t *testing.T) {
	t.Parallel()

	backend := []int{1, 2, 3, 4, 5}
	st := stack.Stack[int](backend)

	for st.Len() > 0 {
		st.Pop()

		assertEq(t, 0, backend[st.Len()], "stack{...}.Pop() should zero out the element")
	}
}

func TestStack_FreeAfterPopMany(t *testing.T) {
	t.Parallel()

	backend := []int{1, 2, 3, 4, 5}
	st := stack.Stack[int](backend)

	for st.Len() > 0 {
		st.PopMany(nil, 2)

		n := len(backend) - st.Len()
		want := make([]int, n)
		assertEqSlices(t, want, backend[st.Len():],
			"stack{len=%d}.PopMany() should zero out the elements", st.Len())
	}
}

func assertEq[E comparable](t *testing.T, want, got E, f string, args ...any) {
	t.Helper()

	if got != want {
		t.Errorf(f, args...)
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertEqSlices[E comparable](t *testing.T, want, got []E, f string, args ...any) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf(f, args...)
		t.Errorf("got %v, want %v", got, want)
	}
}
