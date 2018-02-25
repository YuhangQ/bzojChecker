package json

import (
	"encoding/json"
)

func Prase(data []byte) (map[string]interface{}, error) {
	var f interface{}
	err := json.Unmarshal(data, &f)
	m := f.(map[string]interface{})
	return m, err
}
