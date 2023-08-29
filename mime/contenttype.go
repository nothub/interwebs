package mime

import (
	"github.com/elnormous/contenttype"
)

func Negotiate(header string, provided []string) (string, error) {
	types, err := ParseMediaTypes(provided...)
	if err != nil {
		return "", err
	}
	mt, _, err := contenttype.GetAcceptableMediaTypeFromHeader(header, types)
	if err != nil {
		return "", err
	}
	return mt.MIME(), nil
}

var cache = make(map[string]contenttype.MediaType)

func ParseMediaTypes(inputs ...string) (types []contenttype.MediaType, err error) {
	for _, s := range inputs {
		if mt, ok := cache[s]; ok {
			types = append(types, mt)
			continue
		}
		mt, err := contenttype.ParseMediaType(s)
		if err != nil {
			return nil, err
		}
		types = append(types, mt)
		cache[s] = mt
	}
	return types, nil
}
