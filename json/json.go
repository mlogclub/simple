package json

import (
	"encoding/json"
)

func Parse(str string, t interface{}) error {
	return json.Unmarshal([]byte(str), t)
}

func ToStr(t interface{}) (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
