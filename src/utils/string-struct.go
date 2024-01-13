package utils

import (
	"encoding/json"
)

func StructToString(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func StringToStruct(jsonStr string, target interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), target)
	if err != nil {
		return err
	}
	return nil
}
