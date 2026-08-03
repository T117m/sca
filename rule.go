package main

import (
	"strconv"
	"strings"
	"unicode"
)

func validateRule(rule string) (b, s string, ok bool) {
	hasB, hasF, hasS := false, false, false

	for _, r := range strings.ToLower(rule) {
		if !unicode.IsNumber(r) {
			switch r {
			case 'b':
				if hasB {
					return "", "", false
				}
				hasB = true
			case '/':
				if !hasB {
					return "", "", false
				}
				hasF = true
			case 's':
				if !hasB || !hasF {
					return "", "", false
				}
				hasS = true
			default:
				return "", "", false
			}
		} else if d, err := strconv.Atoi(string(r)); err != nil || !hasB || (hasF && !hasS) || d > 8 {
			return "", "", false
		}
	}

	if !hasB || !hasF || !hasS {
		return "", "", false
	}

	split := strings.Split(rule, "/")
	b = strings.TrimPrefix(split[0], "b")
	s = strings.TrimPrefix(split[1], "s")

	return b, s, true
}

func applyRule(alive bool, n uint8, b, s string) bool {
	if alive && has(s, n) {
		return true
	} else if !alive && has(b, n) {
		return true
	}

	return false
}

func has(s string, x uint8) bool {
	for _, r := range s {
		if n, err := strconv.Atoi(string(r)); err == nil {
			if uint8(n) == x {
				return true
			}
		}
	}

	return false
}
