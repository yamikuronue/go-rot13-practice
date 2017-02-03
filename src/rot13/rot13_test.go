package rot13

import (
	"testing"
)

func TestEncodeReturnsString(t *testing.T) {
	/* table based test case*/
	table := []struct {
		input    string
		expected string
	}{
		/* Easy case: first half of alphabet*/
		{"aabbcc", "nnoopp"},
		/* Harder case: last half of alphabet*/
		{"nnoopp", "aabbcc"},
		/*Capital letters: first half*/
		{"AABBCC", "NNOOPP"},
		/*Capital letters: last half*/
		{"NNOOPP", "AABBCC"},
		/*Non-letters*/
		{"@?_~[{|", "@?_~[{|"},
	}

	for _, test := range table {
		output := EncodeRot13(test.input)
		if output != test.expected {
			t.Errorf("Incorrect string returned for input %s; expected %s but got %s", test.input, test.expected, output)
		}
	}
}
