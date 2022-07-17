package utils

import (
	"encoding/json"
	"io"
)

func GetJsonBody(raw io.ReadCloser) map[string]string {
	data := make(map[string]string)
	err := json.NewDecoder(raw).Decode(&data)
	if err != nil {
		panic(err)
	}
	return data
}