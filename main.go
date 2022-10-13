package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tidwall/gjson"
)

var jsonFile *string
var outputFile *string

func init() {
	jsonFile = flag.String("f", "", "json file")
	outputFile = flag.String("o", "", "output file")
}

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("path must be defined")
	}

	var jsonStr string
	var err error
	if *jsonFile != "" {
		jsonStr, err = openJsonFile(*jsonFile)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		jsonStr = readFromStandardInput()
	}

	var result gjson.Result
	result = selectJson(jsonStr, os.Args[1])

	var outputByte []byte
	outputByte, err = json.Marshal(result.Value())
	if err != nil {
		log.Fatal(err.Error())
	}

	if *outputFile != "" {

	}

	fmt.Println(string(outputByte))
	os.Exit(0)
}

func openJsonFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func selectJson(content, path string) gjson.Result {
	return gjson.Get(content, path)
}

func readFromStandardInput() string {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		str += txt
	}
	return str
}