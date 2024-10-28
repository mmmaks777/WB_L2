package main

import "testing"

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{`qwe\4\5`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{`qwe\`, "", true},         // незавершенный escape
		{`\4abcd`, "4abcd", false}, // некорректная строка
	}

	for _, tc := range testCases {
		result, err := Unpack(tc.input)
		if tc.hasError {
			if err == nil {
				t.Errorf("expected error for input %q, but got none", tc.input)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for input %q: %v", tc.input, err)
			} else if result != tc.expected {
				t.Errorf("expected %q, got %q for input %q", tc.expected, result, tc.input)
			}
		}
	}
}
