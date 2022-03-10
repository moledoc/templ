# templ

package templ // import "github.com/moledoc/templ"

Package templ implements tools to parse a string or file template, where
parameters inside ${} are substituted with given values. Substitution values
must be provided via a `map`.

FUNCTIONS

func ParseFile(infile string, params map[string]string) (parsed string)
    ParseFile is a function, that parses given file template with given
    parameters.

func ParseStr(str string, params map[string]string) (parsed string)
    ParseFile is a function, that parses given string template with given
    parameters.

## Author

Meelis Utt
