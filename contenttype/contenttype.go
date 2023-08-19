package contenttype

import (
	"errors"
	"github.com/nothub/interwebs/mime"
	"strings"
)

var (
	ErrNoCommonMimeType = errors.New("no common mime type")
)

// Negotiate tries to find a common mime type to select a format for presenting content.
// Provided mime types are prioritized by order.
// Parameters are stripped and ignored.
// Invalid mime types are ignored.
// If no common mime type is found, an empty string will be returned.
func Negotiate(header string, provided []string) (mime.Type, error) {
	if len(provided) == 0 {
		return mime.Type{}, ErrNoCommonMimeType
	}

	// client wants one of these mime types
	var clientTypes []mime.Type

	// server provides these mime types
	var serverTypes []mime.Type

	// parse 'Accept' header sent by the client
	for _, s := range strings.Split(header, ",") {
		mt, err := mime.Parse(s)
		if err != nil {
			return mime.Type{}, err
		}
		clientTypes = append(clientTypes, mt)
	}

	// parse mime types provided by the server
	for _, s := range provided {
		mt, err := mime.Parse(s)
		if err != nil {
			return mime.Type{}, err
		}
		serverTypes = append(serverTypes, mt)
	}

	// TODO: sort client types by quality factor

	for _, cmt := range clientTypes {
		for _, smt := range serverTypes {
			var main string
			if cmt.Main == "*" || cmt.Main == smt.Main {
				main = smt.Main
			}
			var sub string
			if cmt.Sub == "*" || cmt.Sub == smt.Sub {
				sub = smt.Sub
			}
			if main != "" && sub != "" {
				return mime.Type{Main: main, Sub: sub, Params: cmt.Params}, nil
			}
		}
	}

	return mime.Type{}, ErrNoCommonMimeType
}
