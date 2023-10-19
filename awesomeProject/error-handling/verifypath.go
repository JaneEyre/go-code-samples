package main

import "regexp"

func isValidPath(p string) bool {
	pathRe := regexp.MustCompile(`(invalid regular expression`)
	return pathRe.MatchString(p)
}
