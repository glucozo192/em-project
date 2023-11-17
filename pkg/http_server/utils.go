package http_server

import (
	"strings"

	"golang.org/x/exp/slices"
)

func isCheckingPassed(Paths []string, inComing string) bool {
	return slices.ContainsFunc(Paths, func(expected string) bool {
		return isMatchPath(inComing, expected)
	})
}

func isMatchPath(actual, expected string) bool {
	splitedInComingPathElems := strings.Split(actual, "/")
	splitedExpectedElems := strings.Split(expected, "/")
	if len(splitedInComingPathElems) != len(splitedExpectedElems) {
		return false
	}
	isValid := true
	for i, expected := range splitedExpectedElems {
		isValid = isValid && (expected == splitedInComingPathElems[i] || expected == "*")
	}

	return isValid
}
