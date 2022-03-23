package stringsSpread

import (
	"fmt"
	"regexp"
	"strings"
)

func HumpToUnderline(s string) string {
	re, _ := regexp.Compile("[A-Z][^A-Z]*")
	fmt.Println(re.FindAllString(s, -1))
	return ""
}

func UnderlineToHump(s string) string {
	return strings.ReplaceAll(
		strings.Title(
			strings.Join(
				strings.Split(s, "_"),
				" ",
			),
		),
		" ", "",
	)
}
