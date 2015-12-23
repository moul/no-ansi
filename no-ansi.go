package noansi

import "regexp"

func NoAnsiString(input string) (string, error) {
	r := regexp.MustCompile("\x1b\\[[>0-9;]*[mc]")
	return r.ReplaceAllString(input, ""), nil
}
