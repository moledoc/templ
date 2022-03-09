package templ

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return string(dat)
}

var findParam *regexp.Regexp = regexp.MustCompile("\\${.*?}")
var cleanParam *regexp.Regexp = regexp.MustCompile("\\$|\\{|\\}")

func ParseFile(infile string, params map[string]string) (parsed string) {
	f, err := os.Open(infile)
	defer f.Close()
	check(err)
	scanner := bufio.NewScanner(f)
	defer func() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	fn := func(s string) string {
		s = cleanParam.ReplaceAllString(s, "")
		val, ok := params[s]
		if !ok {
			panic(fmt.Sprintf("No value given for parameter: ${%v}.\n", s))
		}
		return val
	}
	for scanner.Scan() {
		line := scanner.Text()
		if findParam.MatchString(line) {
			line = findParam.ReplaceAllStringFunc(line, fn)
		}
		parsed += line + "\n"
	}
	return parsed
}

func ParseStr(str string, params map[string]string) (parsed string) {
	fn := func(s string) string {
		s = cleanParam.ReplaceAllString(s, "")
		val, ok := params[s]
		if !ok {
			panic(fmt.Sprintf("No value given for parameter: ${%v}.\n", s))
		}
		return val
	}
	if findParam.MatchString(str) {
		return findParam.ReplaceAllStringFunc(str, fn) + "\n"
	}

	return parsed
}
