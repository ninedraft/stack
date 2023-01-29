package stack

// Stack is a LIFO data structure.
// It is not thread-safe.
// It's backed by a slice, sot it can be used with `slices` package.
//
// To preallocate a stack, use `make`:
//
//	st := make(stack.Stack[int], 0, 10)
type Stack[E any] []E

func (s *Stack[E]) Push(e E) {
	*s = append(*s, e)
}

// PushMany pushes all elements from the given slice to the stack.
// Order of elements is preserved.
//
//	stack{}.PushMany(1, 2, 3) => stack{1, 2, 3}, stack.Peek() == 3
func (s *Stack[E]) PushMany(e ...E) {
	*s = append(*s, e...)
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, the returned element is the zero value of the element type and the second return value is false.
func (s *Stack[E]) Pop() (E, bool) {
	var empty E
	stack := *s
	l := len(stack)

	if l == 0 {
		return empty, false
	}

	i := l - 1
	e := stack[i]
	stack[i] = empty
	*s = stack[:i]

	return e, true
}

// PopMany removes and appends the top n elements from the stack to the given slice.
// If the stack is empty, the returned slice is the same as the given slice.
// If the stack has less than n elements, all elements are removed.
//
// Elements are appended to the given slice in the reverse order.
//
//	stack{1, 2, 3}.PopMany(nil, 2) => []int{3, 2}, stack{1}
func (s *Stack[E]) PopMany(dst []E, n int) []E {
	stack := *s
	l := len(stack)

	if n > l {
		n = l
	}

	tail := stack[l-n:]
	dst = append(dst, tail...)
	reverse(dst[len(dst)-len(tail):])

	empties(tail)

	*s = stack[:l-n]

	return dst
}

// Peek returns the top element from the stack without removing it.
// If the stack is empty, the returned element is the zero value of the element type and the second return value is false.
func (s *Stack[E]) Peek() (E, bool) {
	var empty E
	stack := *s
	if len(stack) == 0 {
		return empty, false
	}
	return stack[len(stack)-1], true
}

// PeekMany appends the top n elements from the stack to the given slice without removing them.
// If the stack is empty, the returned slice is the same as the given slice.
func (s *Stack[E]) PeekMany(dst []E, n int) []E {
	stack := *s
	l := len(stack)

	if n > l {
		n = l
	}

	tail := stack[l-n:]
	dst = append(dst, tail...)
	reverse(dst[len(dst)-len(tail):])

	return dst
}

func (s *Stack[E]) Len() int {
	return len(*s)
}

func empties[E any](dst []E) {
	if len(dst) == 0 {
		return
	}

	var empty E
	dst[0] = empty
	i := 1
	for i < len(dst) {
		copy(dst[i:], dst[:i])
		i *= 2
	}
}

func reverse[E any](s []E) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
