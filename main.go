package main

import (
	"bufio"
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

	// selected path must defined
	if len(os.Args) < 2 {
		log.Fatal("path must be defined")
	}

	var jsonStr string
	var err error
	var result interface{}

	if *jsonFile != "" {
		jsonStr, err = openJsonFile(*jsonFile)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		jsonStr = readFromStandardInput()
	}

	result = selectJson(jsonStr, os.Args[1])
	fmt.Println(result)

	// send to output file if flag 'o' is provided
	if *outputFile != "" {

	}


	// send to standard output

	os.Exit(0)
}

func openJsonFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func selectJson(content, path string) interface{} {
	return gjson.Get(content, path)
}

func readFromStandardInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}