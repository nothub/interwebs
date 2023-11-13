package mids

import (
	"net"
	"net/http"
	"strings"
)

const XForwardedFor = "X-Forwarded-For"
const TrueClientIP = "True-Client-IP"
const XRealIP = "X-Real-IP"

// TrustRealIP tries to set the request remote address to the
// original requesters ip by reading the "X-Forwarded-For",
// "True-Client-IP" or "X-Real-IP" header.
//
// Do not use this middleware if your application is not fronted
// by a well configured reverse proxy!
func TrustRealIP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get(XForwardedFor) != "" {
			// X-Forwarded-For holds a list of ip's separated by commas.
			// First element is the original requester.
			// Last element is the most recent proxy.
			l := strings.Split(r.Header.Get(XForwardedFor), ",")
			if len(l) > 0 {
				if ip := net.ParseIP(l[0]); ip != nil {
					r.RemoteAddr = ip.String()
					goto DONE
				}
			}
		}

		if r.Header.Get(TrueClientIP) != "" {
			if ip := net.ParseIP(r.Header.Get(TrueClientIP)); ip != nil {
				r.RemoteAddr = ip.String()
				goto DONE
			}
		}

		if r.Header.Get(XRealIP) != "" {
			if ip := net.ParseIP(r.Header.Get(XRealIP)); ip != nil {
				r.RemoteAddr = ip.String()
				goto DONE
			}
		}

	DONE:
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// StripRealIP removes common real-ip headers from request headers.
func StripRealIP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("X-Forwarded-For")
		r.Header.Del("True-Client-IP")
		r.Header.Del("X-Real-IP")
		r.Header.Del("X-Client-IP")
		r.Header.Del("X-Host")
		r.Header.Del("X-Originating-IP")
		r.Header.Del("Akamai-Client-IP")
		r.Header.Del("CF-Connecting-IP")
		r.Header.Del("Fastly-Client-IP")
		r.Header.Del("X-Cluster-Client-Ip")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
