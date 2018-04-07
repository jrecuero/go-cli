package tools

// Stoif converts a string array into an interface array.
func Stoif(st []string) []interface{} {
	result := make([]interface{}, len(st))
	for i, v := range st {
		result[i] = v
	}
	return result
}

func Mtoif(m map[string]string) map[string]interface{} {
	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
