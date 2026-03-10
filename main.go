package main

import (
	"fmt"
	"go-reloaded/helper"
	"os"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	sampletext := string(data)
	result := helper.ConvertBases(sampletext)
	result = helper.ConvertCases(result)
	result = helper.Punctuate(result)
	result = helper.Replace(result)
	result = helper.EditText(result)


	os.Create("result.txt")

	file := os.WriteFile("result.txt", []byte(result), 0644)

	if file != nil {
		fmt.Println("Result not saved", file)
		return
	}
}


