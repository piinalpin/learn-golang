package pkg

import "encoding/json"

func TypeConverter[T any](source any, destination *T) (*T, bool) {
	isValid := true

	v, err := json.Marshal(source)
	if err != nil {
		return nil, !isValid
	}
		
	err = json.Unmarshal(v, &destination)
	if err != nil {
		return nil, !isValid
	}

	return destination, isValid
}