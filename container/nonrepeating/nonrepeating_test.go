package main

import (
	"testing"
)

//代码覆盖率
//go test -coverprofile=c.out
//go tool cover -html=c.out  // go tool cover 查看其他查看方式

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

//性能测试 go test -bench .
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s += s
	}
	ant := 8

	b.Logf("len(s)=%d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ant {
			b.Errorf("get %d for input %s; expected %d", ant, s, actual)
		}
	}
}
