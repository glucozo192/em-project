package string_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnakeCase(t *testing.T) {

	// fields, values := FieldMap(Actor{})
	// assert.Contains(t, fields, "id")
	// assert.NotNil(t, values)
	type testCase struct {
		input    string
		expected string
	}
	testCases := []testCase{
		{
			input:    "",
			expected: "",
		},
		{
			input:    "CreatedAt",
			expected: "created_at",
		},
		{
			input:    "createdAt",
			expected: "created_at",
		},
		{
			input:    "pictureURL",
			expected: "picture_url",
		}, {
			input:    "ID",
			expected: "id",
		},
	}
	for _, c := range testCases {
		actual := ToSnakeCase(c.input)
		assert.Equal(t, c.expected, actual)
	}
}
