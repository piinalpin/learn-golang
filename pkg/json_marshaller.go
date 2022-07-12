package pkg

import "encoding/json"

func Unmarshal(value []byte, types any) error {
	return json.Unmarshal([]byte(value), types)
}

func Marshal(value any) ([]byte, error) {
	return json.Marshal(value)
}