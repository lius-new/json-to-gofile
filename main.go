package main

import (
	"github.com/lius-new/json-to-gofile/json"
)

func main() {
	if err := json.ConvertJsonFileToGoFile("./example.json", "test"); err != nil {
		panic(err)
	}
}
