package boolhelper

import "strconv"

func ParseBool(b string) bool {
	is, err := strconv.ParseBool(b)
	if err != nil {
		return false
	}
	return is
}
