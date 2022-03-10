// Package templ implements tools to parse a string or file template,
// where parameters inside ${} are substituted with given values.
// Substitution values must be provided via a `map`.
package templ

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// check is a convenience function, that makes error checking shorter.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// readFile is a function to read given file into a string.
func readFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return string(dat)
}

// findParam is a regexp that that matches ${} style parameters.
var findParam *regexp.Regexp = regexp.MustCompile("\\${.*?}")

// cleanParam is a regexp that that matches characters '$', '{', '}'.
// Used to clean the substituted value.
var cleanParam *regexp.Regexp = regexp.MustCompile("\\$|\\{|\\}")

// ParseFile is a function, that parses given file template with given parameters.
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

// ParseFile is a function, that parses given string template with given parameters.
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
