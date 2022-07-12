package pkg

import "encoding/json"

func TypeConverter[T any](source any) (*T, bool) {
	var result T
	isValid := true

	v, err := json.Marshal(source)
	if err != nil {
		return nil, !isValid
	}
		
	err = json.Unmarshal(v, &result)
	if err != nil {
		return nil, !isValid
	}

	return &result, isValid
}