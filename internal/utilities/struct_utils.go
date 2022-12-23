package utilities

import (
	"encoding/json"
	"errors"
	"reflect"
)

func ToMap(data any) (map[string]interface{}, error) {
	kind := reflect.TypeOf(data).Kind()
	if kind != reflect.Struct {
		return nil, errors.New("Data is not struct!")
	}

	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	structMap := map[string]interface{}{}
	err = json.Unmarshal(dataByte, &structMap)
	if err != nil {
		return nil, err
	}

	return structMap, err
}

func KeyToSnakeCase(data map[string]interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{})
	for k, v := range data {
		formatedKey := ToSnakeCase(k)
		result[formatedKey] = v
	}

	return
}
