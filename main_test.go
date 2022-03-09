package templ_test

import (
	"github.com/moledoc/templ"
	"io/ioutil"
	"testing"
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

func TestParseFile(t *testing.T) {
	var test = map[string]string{
		"param": "1",
		"test":  "lalal",
		"this":  "hah",
		"that":  "hoho",
	}
	expected := readFile("test_cases/expected.txt")
	got := templ.ParseFile("test_cases/input.txt", test)
	if expected != got {
		t.Fatalf("Incorrect output:\nExpected\n%v\nGot\n%v\n", expected, got)
	}
}

func TestParseStr(t *testing.T) {
	input := "This is the ${param} test.  \nThis line does not have param.  \nThis line has another param (${test}) in it.  \nThis is line has two params: ${this} and ${that}.  \nThis line does not have param2."
	expected := "This is the 1 test.  \nThis line does not have param.  \nThis line has another param (lalal) in it.  \nThis is line has two params: hah and hoho.  \nThis line does not have param2.\n"
	var test = map[string]string{
		"param": "1",
		"test":  "lalal",
		"this":  "hah",
		"that":  "hoho",
	}
	got := templ.ParseStr(input, test)
	if expected != got {
		t.Fatalf("Incorrect output:\nExpected\n%v\nGot\n%v\n", expected, got)
	}
}
