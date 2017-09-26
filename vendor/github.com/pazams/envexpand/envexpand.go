package envexpand

import (
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^\${([a-zA-Z][a-zA-Z0-9_]*)}`)

func Expand(template []byte) []byte {
	str := ExpandString(string(template))
	return []byte(str)
}

func ExpandString(template string) string {
	ret := make([]byte, 0)

	for len(template) > 0 {
		i := strings.Index(template, "$")
		if i < 0 {
			break
		}
		ret = append(ret, template[:i]...)
		template = template[i:]
		if len(template) > 1 && template[1] == '$' {
			// Treat $$ as $.
			ret = append(ret, '$')
			template = template[2:]
			continue
		}
		name, rest, ok := extract(template)

		if !ok {
			// Malformed; treat $ as raw text.
			ret = append(ret, '$')
			template = template[1:]
			continue
		}
		template = rest

		env, found := os.LookupEnv(name)

		if found {
			ret = append(ret, env...)
		} else {
			ret = append(ret, "${"+name+"}"...)
		}

	}
	ret = append(ret, template...)
	return string(ret)
}

// extract returns the name from a leading "${name}" in str.
func extract(str string) (name string, rest string, ok bool) {

	result := re.FindStringSubmatchIndex(str)

	if len(result) < 4 {
		return
	}

	start := result[2]
	end := result[3]

	// if no results for capture group
	if start == end {
		return
	}

	ok = true
	name = str[start:end]
	rest = str[end+1:]

	return
}
