package pseudostack

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testOp func() ([]string, error)

func TestPsuedostack(t *testing.T) {
	tests := map[string]struct {
		op            testOp
		expected      interface{}
		expectedError string
	}{
		"push": {
			op:       testOpPush,
			expected: []string{"alpha", "bravo"},
		},
		"pop": {
			op:       testOpPop,
			expected: []string{"bravo"},
		},
		"pop-error": {
			op:            testOpPopError,
			expectedError: "empty stack; nothing to pop",
		},
		"size-push": {
			op:       testOpSizePush,
			expected: []string{"2"},
		},
		"size-pop": {
			op:       testOpSizePop,
			expected: []string{"1"},
		},
		"peek": {
			op:       testOpPeek,
			expected: []string{"bravo"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tc.op()
			if err != nil {
				if err.Error() == tc.expectedError {
					return
				}
				t.Fatalf("%v failed %v", name, err)
			}
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func testOpPush() ([]string, error) {
	ps := psuedostack{}

	ps.push("alpha")
	ps.push("bravo")

	return ps.values, nil
}

func testOpPop() ([]string, error) {
	var ps psuedostack
	ps.values, _ = testOpPush()

	popped, err := ps.pop()
	if err != nil {
		return nil, err
	}

	return []string{popped}, nil
}

func testOpPopError() ([]string, error) {
	ps := psuedostack{}
	ps.push("alpha")

	_, err := ps.pop()
	if err != nil {
		return nil, err
	}

	// this one should error
	_, err = ps.pop()
	if err != nil {
		return nil, err
	}

	return []string{"unacceptable"}, nil
}

func testOpSizePush() ([]string, error) {
	ps := psuedostack{}
	ps.push("alpha")
	ps.push("bravo")

	return []string{strconv.Itoa(ps.size())}, nil
}

func testOpSizePop() ([]string, error) {
	ps := psuedostack{}
	ps.push("alpha")
	ps.push("bravo")

	_, _ = ps.pop()

	return []string{strconv.Itoa(ps.size())}, nil
}

func testOpPeek() ([]string, error) {
	var ps psuedostack
	ps.values, _ = testOpPush()

	peeked, err := ps.peek()
	if err != nil {
		return nil, err
	}

	return []string{peeked}, nil
}
