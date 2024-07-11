package main

import (
	"reflect"
	"testing"
)

func TestParseLines(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected []quiz
	}{
		{
			name: "Typical input",
			input: [][]string{
				{"Question 1", " Answer 1 "},
				{"Question 2", "Answer 2"},
			},
			expected: []quiz{
				{q: "Question 1", a: "Answer 1"},
				{q: "Question 2", a: "Answer 2"},
			},
		},
		{
			name:  "Empty input",
			input: [][]string{},
			expected: []quiz{},
		},
		{
			name: "Input with extra spaces",
			input: [][]string{
				{" Question 1 ", " Answer 1 "},
				{" Question 2 ", "  Answer 2  "},
			},
			expected: []quiz{
				{q: " Question 1 ", a: "Answer 1"},
				{q: " Question 2 ", a: "Answer 2"},
			},
		},
		{
			name: "Single entry",
			input: [][]string{
				{"Single question", "Single answer"},
			},
			expected: []quiz{
				{q: "Single question", a: "Single answer"},
			},
		},
		{
			name: "No answer",
			input: [][]string{
				{"No answer", ""},
			},
			expected: []quiz{
				{q: "No answer", a: ""},
			},
		},
		{
			name: "Whitespace only answer",
			input: [][]string{
				{"Whitespace only", "    "},
			},
			expected: []quiz{
				{q: "Whitespace only", a: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseLines(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseLines(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}