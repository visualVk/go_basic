package main

import (
	"testing"
	"unicode/utf8"
)

func TestChineseSupport(t *testing.T) {

	table := []struct {
		s   string
		len int
	}{
		{"啊啊啊", 3},
		{"不不", 1},
	}

	for _, tt := range table {
		actual := utf8.RuneCountInString(tt.s)

		if actual != tt.len {
			t.Errorf("expect %s's len is %d, but got %d", tt.s, tt.len, actual)
		}
	}
}

func BenchmarkString(b *testing.B) {
	ans := 39
	s := "啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊"

	for i := 0; i < b.N; i++ {
		actual := utf8.RuneCountInString(s)

		if actual != ans {
			b.Errorf("expect %s's len is %d, but got %d", s, ans, actual)
		}
	}
}
