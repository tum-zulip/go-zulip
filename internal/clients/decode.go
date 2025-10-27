package clients

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"regexp"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
)

func decode(v interface{}, b []byte, contentType string) error {
	if len(b) == 0 {
		return nil
	}

	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(**os.File); ok {
		var err error
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return err
		}
		_, err = (*f).Write(b)
		if err != nil {
			return err
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return err
	}
	if xmlCheck.MatchString(contentType) {
		return xml.Unmarshal(b, v)
	}
	if jsonCheck.MatchString(contentType) {
		return json.Unmarshal(b, v)
	}
	return errors.New("undefined Response type")
}
