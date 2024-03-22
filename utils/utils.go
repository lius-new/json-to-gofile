package utils

import (
	"fmt"
	"reflect"
	"sort"
)

// stringsToStructString: []string to struct string
func StringsToStructString(keys []string) string {
	suffix := "type %s struct {\n"

	structName := ""
	for _, key := range keys {
		suffix += "\t" + key + " string\n"
		structName += key
	}

	content := suffix + "}\n"
	return fmt.Sprintf(content, structName)
}

// getMapKeys: get keys with map[string]interface{}
func GetMapKeys(js map[string]interface{}) (keys []string) {
	for key := range js {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// stringsContainerString: []string whether it contains the current string
func StringsContainerString(strs []string, str string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}

// isArray: is []interface {}
func IsArray(value interface{}) bool {
	return reflect.TypeOf(value).String() == "[]interface {}"
}
