package filter

import "strconv"

func ParseStringToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}
