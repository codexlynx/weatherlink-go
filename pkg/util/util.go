package util

import (
	"net/url"
	"strings"
)

func GetKeys(input map[string]string) []string {
	var keys []string
	for key := range input {
		keys = append(keys, key)
	}
	return keys
}

func RemoveCharacters(input string, characters []string) string {
	output := input
	for _, char := range characters {
		output = strings.ReplaceAll(output, char, "")
	}
	return output
}

func RemoveParameters(keys []string, values url.Values) url.Values {
	for _, key := range keys {
		values.Del(key)
	}
	return values
}
