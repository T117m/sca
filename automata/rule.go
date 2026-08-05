package automata

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func ValidateRule(rule string) (b, s []uint8, ok bool) {
	hasB, hasF, hasS := false, false, false

	for _, r := range strings.ToLower(rule) {
		if !unicode.IsNumber(r) {
			switch r {
			case 'b':
				if hasB {
					return nil, nil, false
				}
				hasB = true
			case '/':
				if !hasB {
					return nil, nil, false
				}
				hasF = true
			case 's':
				if !hasB || !hasF {
					return nil, nil, false
				}
				hasS = true
			default:
				return nil, nil, false
			}
		} else if d, err := strconv.Atoi(string(r)); err != nil || !hasB || (hasF && !hasS) || d > 8 {
			return nil, nil, false
		}
	}

	if !hasB || !hasF || !hasS {
		return nil, nil, false
	}

	var (
		split = strings.Split(strings.ToLower(rule), "/")
		bString = strings.TrimPrefix(split[0], "b")
		sString = strings.TrimPrefix(split[1], "s")
	)

	b = make([]uint8, len(bString))
	s = make([]uint8, len(sString))

	for i, r := range bString {
		d, _ := strconv.Atoi(string(r))
		b[i] = uint8(d)
	}

	for i, r := range sString {
		d, _ := strconv.Atoi(string(r))
		s[i] = uint8(d)
	}

	return b, s, true
}

func applyRule(alive bool, n uint8, b, s []uint8) bool {
	if alive && slices.Contains(s, n) {
		return true
	} else if !alive && slices.Contains(b, n) {
		return true
	}

	return false
}

