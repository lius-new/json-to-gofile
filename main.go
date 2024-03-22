package main

import "github.com/lius/json-to-gofile/json"

func main() {
	json.ConvertJsonFileToGoFile("./vcard.json", "test")
}
