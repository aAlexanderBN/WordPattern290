package main

import (
	"testing"
)

var testCases = []struct {
	name     string
	pattern  string
	s        string
	expected bool
}{
	{name: "case 1", pattern: "abba", s: "dog cat cat dog", expected: true},
	{name: "case 2", pattern: "abba", s: "dog cat cat fish", expected: false},
	{name: "case 3", pattern: "aaaa", s: "dog cat cat dog", expected: false},
	{name: "case 4", pattern: "aaaa", s: "dog dog dog dog", expected: true},
	{name: "case 5", pattern: "abba", s: "dog dog dog dog", expected: false},
}

func TestAll(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := wordPattern(tc.pattern, tc.s)
			if tc.expected != result {
				if tc.expected == true {

					t.Errorf("Ожидалось что pattern %v состоит из s %v ", tc.pattern, tc.s)

				} else {
					t.Errorf("Ожидалось что pattern %v состоит из s %v ", tc.pattern, tc.s)
				}

			}
		})

	}

}
