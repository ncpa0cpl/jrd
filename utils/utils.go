package utils

import "strings"

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func Trim(string []byte) []byte {
	for {
		if isWhitespace(string[0]) {
			string = string[1:]
		} else {
			break
		}
	}

	for {
		if isWhitespace(string[len(string)-1]) {
			string = string[:len(string)-1]
		} else {
			break
		}
	}

	return string
}

func SliceCmp(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func SplitKey(key string) (string, string) {
	dotIdx := strings.Index(key, ".")

	if dotIdx == -1 {
		return key, ""
	}

	return key[:dotIdx], key[dotIdx+1:]
}

func TypeOf(value interface{}) string {
	switch value.(type) {
	case []interface{}:
		return "array"
	case map[string]interface{}:
		return "object"
	case string:
		return "string"
	case float64:
		return "number"
	case bool:
		return "bool"
	case nil:
		return "null"
	}

	return "null"
}
