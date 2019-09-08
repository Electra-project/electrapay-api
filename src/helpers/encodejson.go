package helpers

import "encoding/json"

func EncodeJson(v interface{}) ([]byte, error) {
	data, err := json.Marshal(&v)
	if err != nil {
		return nil, err
	}
	return data, nil
}
