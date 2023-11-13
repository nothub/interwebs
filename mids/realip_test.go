package mids

import (
	"net/http"
	"testing"
)

type h struct{}

func (_ h) ServeHTTP(http.ResponseWriter, *http.Request) {}

func TestTrustRealIP(t *testing.T) {
	type data struct {
		headers    http.Header
		remoteAddr string
	}
	type testCase struct {
		name     string
		inputs   data
		expected string
	}
	var cases []testCase

	cases = append(cases, testCase{
		name: "no headers",
		inputs: data{
			remoteAddr: "192.0.2.1",
		},
		expected: "192.0.2.1",
	})

	headers := make(http.Header)
	headers.Set(XRealIP, "203.0.113.1")
	headers.Set(TrueClientIP, "198.51.100.1")
	cases = append(cases, testCase{
		name: "trust True-Client-IP",
		inputs: data{
			headers:    headers,
			remoteAddr: "192.0.2.1",
		},
		expected: "198.51.100.1",
	})

	headers = make(http.Header)
	headers.Set(XRealIP, "203.0.113.1")
	headers.Set(XForwardedFor, "198.51.100.1,198.51.100.2,198.51.100.3")
	headers.Set(TrueClientIP, "203.0.113.2")
	cases = append(cases, testCase{
		name: "trust X-Forwarded-For",
		inputs: data{
			headers:    headers,
			remoteAddr: "192.0.2.1",
		},
		expected: "198.51.100.1",
	})

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := http.Request{}
			req.Header = tc.inputs.headers
			req.RemoteAddr = tc.inputs.remoteAddr
			TrustRealIP(h{}).ServeHTTP(nil, &req)
			if req.RemoteAddr != tc.expected {
				t.Fatalf("remote address was wrong!\nactual:   %++v\nexpected: %++v\n", req.RemoteAddr, tc.expected)
			}
		})
	}
}

func TestStripRealIP(t *testing.T) {
	req := http.Request{}
	headers := make(http.Header)
	headers.Set("foo", "bar")
	headers.Set("True-Client-IP", "198.51.100.1")
	headers.Set("CF-Connecting-IP", "198.51.100.2")
	req.Header = headers

	StripRealIP(h{}).ServeHTTP(nil, &req)
	if val := req.Header.Get("foo"); val != "bar" {
		t.Fail()
		t.Logf("header %q should have value %q but has value %q\n", "foo", "bar", val)
	}
	if req.Header.Get("True-Client-IP") != "" {
		t.Fail()
		t.Logf("header %q should not have any value set\n", "True-Client-IP")
	}
	if req.Header.Get("CF-Connecting-IP") != "" {
		t.Fail()
		t.Logf("header %q should not have any value set\n", "CF-Connecting-IP")
	}
}
