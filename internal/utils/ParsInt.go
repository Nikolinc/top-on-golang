package utils

import "strconv"

func ParsInt16(str string) int16 {
	v, _ := strconv.ParseInt(str, 10, 16)
	return int16(v)
}

func ParsInt32(str string) int32 {
	v, _ := strconv.ParseInt(str, 10, 32)
	return int32(v)
}
