package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput []string
	}{
		{
			input:          "This is a Test String",
			expectedOutput: []string{"this", "is", "a", "test", "string"},
		},
		{
			input:          "Another Test String",
			expectedOutput: []string{"another", "test", "string"},
		},
		{
			input:          "ALL CAPS",
			expectedOutput: []string{"all", "caps"},
		},
		{
			input:          "1234",
			expectedOutput: []string{"1234"},
		},
		{
			input:          "",
			expectedOutput: []string{},
		},
	}

	for _, tc := range testCases {
		output := cleanInput(tc.input)

		if !reflect.DeepEqual(output, tc.expectedOutput) {
			t.Errorf("cleanInput(%q) = %q, want %q", tc.input, output, tc.expectedOutput)
		}
	}
}
