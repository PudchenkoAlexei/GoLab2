package Lab_2

import (
	"bytes"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type ComputeSuite struct{}

var _ = check.Suite(&ComputeSuite{})

func (s *ComputeSuite) TestCompute(c *check.C) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr bool
	}{
		{
			name:        "Valid Input",
			input:       "+ * 1 2 3",
			expected:    "1 2 * 3 +",
			expectedErr: false,
		},
		{
			name:        "Invalid Expression",
			input:       "+ 5 *",
			expected:    "",
			expectedErr: true,
		},
		{
			name:        "Invalid Input",
			input:       "+$ 5 4",
			expected:    "",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		c.Log("Test Case: ", test.name)
		input := bytes.NewBufferString(test.input)
		output := &bytes.Buffer{}
		handler := ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()

		if test.expectedErr {
			c.Check(err, check.NotNil)
		} else {
			c.Check(err, check.IsNil)
			c.Check(output.String(), check.Equals, test.expected)
		}
	}
}
