package utility

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadLines(path string) []string {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func ReadLinesToBytes(path string) []byte {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return fileContent
}
