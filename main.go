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

		// if the log entry does not start with a { character
		// then it's a entry that looks like :-
		// 2020/04/20 06:45:59.015 info http.log.access.log0
		// then extract the json object from the entry
		if string(input[0]) != "{" && strings.Contains(input, "handled request") == true {
			input = input[strings.Index(input, "handled request ")+15 : len(input)]
			input = strings.TrimSpace(input)
		}

		// if the log entry starts with a { character
		// then it must be a complete JSON object like :-
		// {"level":"info","ts":1588143602.7486432,"logger":"http.log.access.log0",
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
	os.Exit(0)
}

func removeSquareBrackets(input string) string {
	input = strings.Replace(input, "[", "", 1)
	input = strings.Replace(input, "]", "", 1)
	return input
}
