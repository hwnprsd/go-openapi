package types

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func getType(data interface{}) map[string]MediaType {
	result := make(map[string]MediaType)
	schema := createSchema(data)
	mediaType := MediaType{
		Schema: schema,
	}
	result["application/json"] = mediaType
	return result
}

func createSchema(data interface{}) Schema {
	fmt.Println("data", data)
	var schema Schema

	dataType := reflect.TypeOf(data)
	// For some reason, nil = string
	if dataType == nil {
		return Schema{
			Type: "string",
		}
	}

	if dataType.Kind() == reflect.Slice {
		elemType := dataType.Elem()
		elemSchema := createSchema(reflect.Zero(elemType).Interface())
		schema = Schema{
			Type:  "array",
			Items: &elemSchema,
		}
		return schema
	}

	if dataType.Kind() == reflect.Map || dataType.Kind() == reflect.Struct {
		properties := make(map[string]Schema)

		// Marshal and unmarshal the data to a map
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			// Handle error
			return schema
		}

		var obj map[string]interface{}
		if err := json.Unmarshal(jsonBytes, &obj); err != nil {
			// Handle error
			return schema
		}

		for k, v := range obj {
			properties[k] = createSchema(v)
		}

		schema = Schema{
			Type:       "object",
			Properties: properties,
		}
	} else {
		// Map Go primitive types to OpenAPI types
		switch dataType.Kind() {
		case reflect.String:
			schema.Type = "string"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			schema.Type = "integer"
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			schema.Type = "integer"
		case reflect.Float32, reflect.Float64:
			schema.Type = "number"
		case reflect.Bool:
			schema.Type = "boolean"
		default:
			schema.Type = "object" // Fallback type
		}
	}

	return schema
}
