package main

import (
	"regexp"
	"strings"
)

func checkRoom(s string) ([]string, bool) {
	spaces := regexp.MustCompile(`\s+`)
	roomExpression := regexp.MustCompile(`\s*.\s*\d+\s*\d+\s*`)
	roomFlag := roomExpression.MatchString(s)

	if !roomFlag {
		return nil, false
	}
	s = strings.TrimSpace(s)

	if s[0] == 'L' || s[0] == 'l' || s[0] == '#' {
		return nil, false
	}

	lines := spaces.Split(s, 3)

	for _, s := range lines {
		if s == "" || s == " " {
			return nil, false
		}
	}
	return lines, true
}

func checkPath(s string) ([]string, bool) {
	seperator := regexp.MustCompile(`\s*-\s*`)
	pathExpression := regexp.MustCompile(`\s*.\s*-\s*.\s*`)
	pathFlag := pathExpression.MatchString(s)

	if !pathFlag {
		return nil, false
	}
	s = strings.TrimSpace(s)
	lines := seperator.Split(s, 2)

	for _, s := range lines {
		if s == "" || s == " " {
			return nil, false
		}
	}

	return lines, true
}
