package Lab_2

import (
	"bytes"
	"testing"
)

func TestCompute(t *testing.T) {
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
			expected:    "invalid expression",
			expectedErr: true,
		},
		{
			name:        "Invalid Input",
			input:       "+$ 5 4",
			expected:    "invalid input",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := bytes.NewBufferString(test.input)
			output := &bytes.Buffer{}
			handler := ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()

			if (err != nil) != test.expectedErr {
				t.Errorf("expected error: %v, actual: %v", test.expectedErr, err)
			}

			if err != nil && err.Error() != test.expected {
				t.Errorf("expected error message: %s, actual: %s", test.expected, err.Error())
			}

			if err == nil && output.String() != test.expected {
				t.Errorf("expected output: %s, actual: %s", test.expected, output.String())
			}
		})
	}
}
