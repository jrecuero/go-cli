package tools

import (
	"reflect"
)

// Stoif converts a string array into an interface array.
func Stoif(st []string) []interface{} {
	result := make([]interface{}, len(st))
	for i, v := range st {
		result[i] = v
	}
	return result
}

// Mtoif converts a string map string into an interface array.
func Mtoif(m map[string]string) map[string]interface{} {
	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// GetReflectType returns a string with the type for the variable.
func GetReflectType(v interface{}) string {
	if v == nil {
		return "<nil-type>"
	}
	r := reflect.ValueOf(v)
	return r.Type().String()
}

// MapCast returns a map entry properly casted.
func MapCast(value interface{}) map[string]interface{} {
	return value.(map[string]interface{})
}

// KeysForMap returns all keys for a map as a  list
func KeysForMap(table map[string]interface{}) []string {
	keys := make([]string, len(table))
	i := 0
	for k := range table {
		keys[i] = k
		i++
	}
	return keys
}
