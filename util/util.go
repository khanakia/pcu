package util

import (
	"fmt"
	"regexp"
	"strings"
)

/*
 * v6.22.0 > ^6.22
 * 3.7.1 > ^3.7
 * 1.6.x > ^1.6
 */
func TagNameConvert(name string) string {
	re := regexp.MustCompile("[0-9].+")
	result := re.FindString(name)

	arr := strings.Split(result, ".")

	if len(arr) > 1 {
		return fmt.Sprintf("^%s.%s", arr[0], arr[1])
	}

	return ""
}
