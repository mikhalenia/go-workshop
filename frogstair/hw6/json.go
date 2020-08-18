package hw6

import "encoding/json"

// Decode decodes an interface into JSON
func Decode(jsonStr string, i interface{}) (interface{}, error) {
	err := json.Unmarshal([]byte(jsonStr), &i)
	return i, err
}

// Encode encodes an interface into JSON
func Encode(i interface{}) (string, error) {
	s, err := json.Marshal(i)
	return string(s), err
}
