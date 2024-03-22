package Lab_2

import (
	"fmt"
	"testing"

	"gopkg.in/check.v1"
)

func TestPrefixToPostfix(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&PrefixToPostfixSuite{})

func (s *PrefixToPostfixSuite) TestPrefixToPostfix(c *check.C) {
	tests := []struct {
		prefix   string
		expected string
	}{
		{"+ 2 3", "2 3 +"},
		{"* + 2 3 4", "2 3 + 4 *"},
		{"+ * 1 2 3", "1 2 * 3 +"},
		{"- 7 + 6 / 5 * 4 - 3 + 1 2", "7 6 5 4 3 1 2 + - * / + -"},
		{"+ 5 * 3 - 4 / 2 + 2 - 3 + 4 2", "5 3 4 2 2 3 4 2 + - + / - * +"},
		{"* 5 + 4 + 1 - 4 + 8 * 4 + 2 ^ 2 5", "5 4 1 4 8 4 2 2 5 ^ + * + - + + *"},
	}

	for _, test := range tests {
		result, err := PrefixToPostfix(test.prefix)
		if test.expected == "" {
			c.Assert(err, check.NotNil, check.Commentf("Expected error for input: %s", test.prefix))
		} else {
			c.Assert(err, check.IsNil, check.Commentf("Unexpected error for input: %s", test.prefix))
			c.Assert(result, check.Equals, test.expected, check.Commentf("Incorrect result for input: %s", test.prefix))
		}
	}
}

func (s *PrefixToPostfixSuite) TestPrefixToPostfixInvalidInput(c *check.C) {
	tests := []struct {
		prefix   string
		expected error
	}{
		{" 2 & + 3", fmt.Errorf("invalid input")},
		{" ", fmt.Errorf("invalid input")},
		{"%", fmt.Errorf("invalid input")},
		{"", fmt.Errorf("invalid input")},
		{"+ 2", fmt.Errorf("invalid expression")},
		{"-+ 2 3", fmt.Errorf("invalid input")},
	}

	for _, test := range tests {
		_, err := PrefixToPostfix(test.prefix)
		c.Assert(err, check.NotNil)
		c.Check(err, check.ErrorMatches, test.expected.Error())
	}
}

func (s *PrefixToPostfixSuite) ExamplePrefixToPostfix(c *check.C) {
	res, _ := PrefixToPostfix("- 5 * 4 2")
	fmt.Println(res)

	// Output:
	// 5 4 2 * -
}
