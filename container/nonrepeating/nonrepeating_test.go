package main

import "testing"

func TestSubstr(t *testing.T) {

	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},

		{"", 0},
		{"b", 1},
		{"bbbbb", 1},
		{"abcabcabcb", 3},

		{"这里是我慕课网", 7},
		{"一二三一二", 3},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if tt.ans != actual {
			t.Errorf("get %d for input %s; expected %d", tt.ans, tt.s, actual)
		}
	}
}
