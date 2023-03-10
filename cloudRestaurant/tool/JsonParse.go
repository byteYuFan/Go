package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {
}

func Decode(io io.ReadCloser, v any) error {
	return json.NewDecoder(io).Decode(v)
}