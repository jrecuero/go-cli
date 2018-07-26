package tools

import (
	"errors"
	"reflect"
	"strings"
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

// SearchKeyInTable checks if the key is in the given array.
func SearchKeyInTable(table []string, key string) error {
	for _, v := range table {
		if key == v {
			return nil
		}
	}
	return errors.New("not found")
}

// SearchKeyInRuneTable checks if the key is in the given array.
func SearchKeyInRuneTable(table []rune, key rune) error {
	for _, v := range table {
		if key == v {
			return nil
		}
	}
	return errors.New("not found")
}

// SearchKeyInMap checks if the key is in the given map.
func SearchKeyInMap(table map[string]interface{}, key string) error {
	for k := range table {
		if key == k {
			return nil
		}
	}
	return errors.New("not found")
}

// GetAllEntriesFromMap returns all remaining entries in the given map.
func GetAllEntriesFromMap(table map[string]interface{}) ([]interface{}, error) {
	var results []interface{}
	for _, entry := range table {
		switch v := entry.(type) {
		case map[string]interface{}:
			locals, err := GetAllEntriesFromMap(v)
			if err != nil {
				return nil, err
			}
			results = append(results, locals...)
		default:
			results = append(results, v)
		}
	}
	return results, nil
}

// SearchSequenceInMap returns all entries that match the sequence.
func SearchSequenceInMap(table map[string]interface{}, sequence []string) ([]interface{}, error) {
	var results []interface{}
	if len(sequence) == 0 {
		locals, err := GetAllEntriesFromMap(table)
		if err != nil {
			return nil, err
		}
		results = append(results, locals...)
		return results, nil
	}
	token := sequence[0]
	entry := table[token]
	if entry != nil {
		switch v := entry.(type) {
		case map[string]interface{}:
			locals, err := SearchSequenceInMap(v, sequence[1:])
			if err != nil {
				return nil, err
			}
			results = append(results, locals...)
		default:
			results = append(results, v)
		}
		return results, nil
	}
	return nil, errors.New("not found")
}

// SearchPatternInMap searches for the given pattern in the commands map.
func SearchPatternInMap(table map[string]interface{}, pattern string) ([]interface{}, error) {
	sequence := strings.Split(pattern, " ")
	return SearchSequenceInMap(table, sequence)
}

// PString returns a pointer to the string.
func PString(st string) *string {
	return &st
}

// String returns the string value contained in a pointer to string.
func String(pst *string) string {
	if pst != nil {
		return *pst
	}
	return ""
}

// ToString returns the string value contained in a pointer to string.
func ToString(in interface{}) string {
	if in != nil {
		return in.(string)
	}
	return ""
}

// LastChar returns the last character in a string,
func LastChar(st string) string {
	return string(st[len(st)-1])
}

// GetValFromArgs returns the value for the given name from arguments being
// passed to any command callback.
func GetValFromArgs(arguments interface{}, name string) interface{} {
	return arguments.(map[string]interface{})[name]
}

// GetStringFromArgs returns the string value for the given name from arguments
// being passed to any command callback.
func GetStringFromArgs(arguments interface{}, name string) string {
	return arguments.(map[string]interface{})[name].(string)
}

// GetIntFromArgs returns the int value for the given name from arguments
// being passed to any command callback.
func GetIntFromArgs(arguments interface{}, name string) int {
	return arguments.(map[string]interface{})[name].(int)
}
