package utils

import "encoding/json"

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	if data != nil {
		jsonData, _ := json.Marshal(data)
		_ = json.Unmarshal(jsonData, &result)
	}
	return result
}

func MapToStruct(m map[string]interface{}, result interface{}) error {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, result)
	return err
}
