package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

var jsonFile *string
var outputFile *string
var prettyPrint *bool

func init() {
	jsonFile = flag.String("f", "", "json file")
	outputFile = flag.String("o", "", "output file")
	prettyPrint = flag.Bool("p", false, "pretty print")
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatalln("path must be defined")
	}

	var jsonStr string
	var err error
	if *jsonFile != "" {
		jsonStr, err = openJsonFile(*jsonFile)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		jsonStr = readFromStandardInput()
	}

	var result *gjson.Result
	result = selectJson(jsonStr, flag.Arg(0))
	var outputByte []byte
	outputByte, err = marshalSelectedJson(result, *prettyPrint)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if *outputFile != "" {
		err = writeOuputToFile(*outputFile, outputByte)
		if err != nil {
			log.Fatalln(err.Error())
		}
		os.Exit(0)
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

func selectJson(content, path string) *gjson.Result {
	res := gjson.Get(content, path)
	return &res
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

func writeOuputToFile(path string, out []byte) (err error) {
	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, 754)
	if err != nil {
		return err
	}
	var f *os.File
	f, err =os.Create(path)

	_, err = f.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func marshalSelectedJson(res *gjson.Result, pretty bool) (out []byte, err error) {
	if pretty {
		return json.MarshalIndent(res.Value(), "", "\t")
	}
	return json.Marshal(res.Value())
}