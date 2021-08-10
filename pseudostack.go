package pseudostack

import (
	"fmt"
)

// Not quite a stack.
// Therefore, this class is designed to handle string slices with the convenience of a stack-like struct.
// Testing goes so far as intended use; beyond that I'd find something else that's (probably a lot) better.
type Psuedostack struct {
	values []string
}

func (ps *Psuedostack) Push(value string) {
	ps.values = append(ps.values, value)
}

func (ps *Psuedostack) Pop() (string, error) {
	n := len(ps.values)
	if n > 0 {
		value := ps.values[n-1]
		ps.values = ps.values[:n-1]
		return value, nil
	}

	return "", fmt.Errorf("empty stack; nothing to pop")
}

func (ps *Psuedostack) Size() int {
	return len(ps.values)
}

func (ps *Psuedostack) Peek() (string, error) {
	n := len(ps.values)
	if n > 0 {
		return ps.values[n-1], nil
	}

	return "", fmt.Errorf("empty stack; nothing to peek")
}
