package stringpkg

import "strconv"

func Str2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
