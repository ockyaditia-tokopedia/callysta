package main

import (
	"regexp"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func BenchmarkRegex(b *testing.B) {
	t := "1999-06-05"
	for i := 0; i < b.N; i++ {
		regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", t)
	}
}

func BenchmarkParse(b *testing.B) {
	t := "1999-06-05"
	format := "2006-01-02"
	for i := 0; i < b.N; i++ {
		time.Parse(format, t)
	}
}
