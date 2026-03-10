package helper

import "regexp"

func Punctuate(s string) string {

	re := regexp.MustCompile(`\s+([.,!?:;]+)`)
	s = re.ReplaceAllString(s, "$1")

	re2 := regexp.MustCompile(`([.,!?:;]+)(\w)`)
	s = re2.ReplaceAllString(s, "$1 $2")

	return s
}
