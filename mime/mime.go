package mime

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidMimeType  = errors.New("invalid mime type")
	ErrInvalidMimeParam = errors.New("invalid mime param")
)

type Params map[string]string

func (p Params) isAscii() bool {
	return strings.ToUpper(p["charset"]) == "US-ASCII"
}

func (p Params) isUtf8() bool {
	return strings.ToUpper(p["charset"]) == "UTF-8"
}

func (p Params) quality() (q float64, err error) {
	q, err = strconv.ParseFloat(p["q"], 64)
	if err != nil {
		return q, err
	}
	return q, nil
}

const (
	Css   = "text/css"
	Csv   = "text/csv"
	Html  = "text/html"
	Js    = "text/javascript"
	Plain = "text/plain"

	Json   = "application/json"
	JsonLd = "application/ld+json"
	Pdf    = "application/pdf"

	Avif      = "image/avif"
	Bmp       = "image/bmp"
	Gif       = "image/gif"
	Jpeg      = "image/jpeg"
	Png       = "image/png"
	Svg       = "image/svg+xml"
	WebmImage = "image/webm"

	Mp3       = "audio/mpeg"
	OggAudio  = "audio/ogg"
	Wav       = "audio/wav"
	WebmAudio = "audio/webm"

	Mp4       = "video/mp4"
	Mpeg      = "video/mpeg"
	OggVideo  = "video/ogg"
	WebmVideo = "video/webm"
)

var specials = []rune{'!', '#', '$', '%', '&', '\'', '*', '+', '-', '.', '^', '_', '`', '|', '~'}

func Parse(raw string) (mime string, params Params, err error) {
	// TODO: verify with https://mimesniff.spec.whatwg.org/#parsing-a-mime-type

	params = NewParams()

	for i, s := range strings.Split(raw, ";") {
		if i == 0 {
			raw = s
			continue
		}
		k, v, err := parseParam(s)
		if err != nil {
			return "", Params{}, err
		}
		params[k] = v
	}

	main, sub, err := parseType(raw)
	if err != nil {
		return "", Params{}, err
	}

	return main + "/" + sub, params, nil
}

func NewParams() (params Params) {
	params = make(map[string]string)
	params["charset"] = "US-ASCII"
	params["q"] = "1.0"
	return params
}

func parseType(s string) (main string, sub string, err error) {
	if !strings.Contains(s, "/") {
		return main, sub, errors.Join(ErrInvalidMimeType, fmt.Errorf("%s", s))
	}

	split := strings.SplitN(s, "/", 2)
	main = split[0]
	sub = split[1]

	s = strings.ToLower(s)

	// trim whitespace
	main = strings.TrimSpace(main)
	sub = strings.TrimSpace(sub)

	for _, r := range main + sub {
		if unicode.IsSpace(r) {
			return "", "", ErrInvalidMimeType
		}
	}
	// TODO: validate contained tokens

	switch main {
	case "":
		return "", "", ErrInvalidMimeType
	case "*":
		switch sub {
		case "":
			return "", "", ErrInvalidMimeType
		case "*":
			return main, sub, nil
		default:
			return "", "", ErrInvalidMimeType
		}
	default:
		switch sub {
		case "":
			return "", "", ErrInvalidMimeType
		case "*":
			return main, sub, nil
		default:
			return main, sub, nil
		}
	}
}

func parseParam(s string) (k string, v string, err error) {
	if !strings.Contains(s, "=") {
		return k, v, errors.Join(ErrInvalidMimeParam, fmt.Errorf("%s", s))
	}

	split := strings.SplitN(s, "=", 2)
	k = split[0]
	v = split[1]

	// key is case-insensitive (but value is not)
	k = strings.ToLower(k)

	// trim whitespace
	k = strings.TrimSpace(k)
	v = strings.TrimSpace(v)

	for _, r := range k + v {
		// TODO: also do this for quoted values?
		if unicode.IsSpace(r) {
			return "", "", ErrInvalidMimeParam
		}
	}
	// TODO: validate contained tokens

	return k, v, nil
}
