package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {
	var common string
	var input string
	var referer string
	var userAgent string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		common = ""
		referer = `""`
		userAgent = `""`

		input = scanner.Text()

		// replace tabs with spaces
		input = strings.Replace(input, "\t", " ", -1)

		// if the input line contains a handled request
		// then extract the json object from the entry
		if strings.Contains(input, "handled request") == true {
			input = input[strings.Index(input, "handled request ")+15 : len(input)]
			input = strings.TrimSpace(input)

			// extract the data elements from the json object
			common = gjson.Get(input, "common_log").String()
			referer = gjson.Get(input, "request.headers.Referer").String()
			userAgent = gjson.Get(input, "request.headers.User-Agent").String()

			// remove square brackets
			referer = removeSquareBrackets(referer)
			userAgent = removeSquareBrackets(userAgent)

			// output combined log format
			fmt.Println(common + " " + referer + " " + userAgent)
		}
	}
	os.Exit(0)
}

func removeSquareBrackets(input string) string {
	input = strings.Replace(input, "[", "", 1)
	input = strings.Replace(input, "]", "", 1)
	return input
}
