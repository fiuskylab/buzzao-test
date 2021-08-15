package utils

import "strconv"

// StrToInt convert string to integer
func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
