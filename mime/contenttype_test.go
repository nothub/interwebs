package mime

import (
	"testing"
)

const firefoxHeader = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"

type test struct {
	name     string
	header   string
	provided []string
	expected string
}

func Test(t *testing.T) {
	tests := []test{
		{
			name:     "simple 1",
			header:   firefoxHeader,
			provided: []string{"text/html"},
			expected: "text/html",
		},
		{
			name:     "simple 2",
			header:   firefoxHeader,
			provided: []string{"application/xml"},
			expected: "application/xml",
		},
		{
			name:     "first of multiple provided",
			header:   firefoxHeader,
			provided: []string{"text/html", "application/xhtml+xml"},
			expected: "text/html",
		},
		{
			name:     "second of multiple provided",
			header:   firefoxHeader,
			provided: []string{"application/xhtml+xml", "text/html"},
			expected: "text/html",
		},
		{
			name:     "wildcard 1",
			header:   firefoxHeader,
			provided: []string{"application/json"},
			expected: "application/json",
		},
		{
			name:     "wildcard 3",
			header:   "text/a;q=0.9,*/*;q=1.0",
			provided: []string{"text/a", "text/b"},
			expected: "text/b",
		},
		{
			name:     "quality priority 1",
			header:   firefoxHeader,
			provided: []string{"application/xml", "image/avif"},
			expected: "image/avif",
		},
		{
			name:     "quality priority 2",
			header:   firefoxHeader,
			provided: []string{"application/xml", "application/json"},
			expected: "application/xml",
		},
		{
			name:     "quality priority 3",
			header:   "text/a;q=0.9,text/b;q=1.0",
			provided: []string{"text/a", "text/b"},
			expected: "text/b",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mt, err := Negotiate(firefoxHeader, test.provided)
			if err != nil {
				t.Logf("error=%++q\n", err)
				t.Fail()
			}
			if mt != test.expected {
				t.Fail()
			}
			if t.Failed() {
				t.Logf("header:   %++q\n", test.header)
				t.Logf("provided: %++q\n", test.provided)
				t.Logf("expected: %++q\n", test.expected)
				t.Logf("actual:   %++q\n", mt)
			}
		})
	}
}
