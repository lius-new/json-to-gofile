package json

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lius-new/json-to-gofile/utils"
	"github.com/lius-new/json-to-gofile/write"
)

// JsonFileToMap: json file to map[string]interface{}
// example:
//
// ```
// jsonMap, err = jsonFileToMap(filePath)
//
// ```
func ConvertJsonFileToMap(filePath string) (content map[string]interface{}, err error) {
	var buf []byte
	if buf, err = os.ReadFile(filePath); err != nil {
		return nil, err
	}

	content = make(map[string]interface{})

	if err = json.Unmarshal(buf, &content); err != nil {
		return nil, err
	}

	return
}

// JsonMapToStrings: map[string]interface{} to []string
func ConvertToJsonMapString(js map[string]interface{}, res *[]string) {
	structString := utils.StringsToStructString(utils.GetMapKeys(js))

	if !utils.StringsContainerString(*res, structString) {
		*res = append(*res, structString)
	}

	for _, v := range js {
		if utils.IsArray(v) {
			for _, innerV := range v.([]interface{}) {
				ConvertToJsonMapString(innerV.(map[string]interface{}), res)
			}
		}
	}
}

// ConvertJsonFileToStructStrings: json file to []string with struct
func ConvertJsonFileToStructStrings(filePath string) (structStrings []string, err error) {
	var jsonMap map[string]interface{}
	if jsonMap, err = ConvertJsonFileToMap(filePath); err != nil {
		return nil, err
	}

	ConvertToJsonMapString(jsonMap, &structStrings)

	return
}

// ConvertJsonFileToFiles: json file to golang file
func ConvertJsonFileToGoFile(jsonFilePath string, goPackageName string) (err error) {
	var structStrings []string

	if structStrings, err = ConvertJsonFileToStructStrings(jsonFilePath); err != nil {
		return err
	}

	newFilePath := strings.TrimRight(jsonFilePath, ".json") + ".go"

	write.WriteContentToFile(newFilePath, fmt.Sprintf("package %s \n", goPackageName))

	for _, v := range structStrings {
		if err = write.WriteContentToFile(newFilePath, v); err != nil {
			break
		}
	}
	return
}
