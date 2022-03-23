package validator

import (
	"fmt"
	"regexp"
	"strings"
)

// 查找连续的小写字母
func FindConsecutiveLowerLetters(text string, n int) []string {
	reg := regexp.MustCompile(`[a-z]+`)
	return reg.FindAllString(text, -1)
}

// 查找连续的非小写字母
func FindConsecutiveNoLowerLetters(text string, n int) []string {
	reg := regexp.MustCompile(`[^a-z]+`)
	return reg.FindAllString(text, -1)
}

// 查找连续的单词字母
func FindWords(text string, n int) []string {
	reg := regexp.MustCompile(`[\w]+`)
	return reg.FindAllString(text, -1)
}

// 查找连续的大写字母
func FindConsecutiveUpperLetters(text string, n int) []string {
	reg := regexp.MustCompile(`[[:upper:]]+`)
	return reg.FindAllString(text, -1)
}

// 查找连续的非 ASCII 字符
func FindConsecutiveNoAscii(text string, n int) []string {
	reg := regexp.MustCompile(`[[:^ascii:]]+`)
	return reg.FindAllString(text, -1)
}

// 查找连续的标点符号
func FindConsecutivePunctuation(text string, n int) []string {
	reg := regexp.MustCompile(`[\pP]+`)
	return reg.FindAllString(text, -1)
}

// 查找连续的非标点符号字符
func FindConsecutiveNoPunctuation(text string, n int) []string {
	reg := regexp.MustCompile(`[\PP]+`)
	return reg.FindAllString(text, -1)
}

// 查找连续的汉字
func FindConsecutiveChinese(text string, n int) []string {
	reg := regexp.MustCompile(`[\p{Han}]+`)
	return reg.FindAllString(text, -1)
}

// 查找连续的非汉字字符
func FindConsecutiveNoChinese(text string, n int) []string {
	reg := regexp.MustCompile(`[\P{Han}]+`)
	return reg.FindAllString(text, -1)
}

// 查找 Hello 或 Go
func FindMulString(text string, search []string, n int) []string {
	reg := regexp.MustCompile(strings.Join(search, "|"))
	return reg.FindAllString(text, -1)
}

// 查找行首以 H 开头，以空格结尾的字符串
func FindStartEndWith(text, start, end string, n int) []string {
	reg := regexp.MustCompile(fmt.Sprintf(`^%s.*%s`, start, end))
	return reg.FindAllString(text, -1)
}

// 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
//reg = regexp.MustCompile(`(?U)^H.*\s`)
//fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello "]

// 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
//reg = regexp.MustCompile(`(?i:^hello).*Go`)
//fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello 世界！123 Go"]

// 查找 Go.
//reg = regexp.MustCompile(`\QGo.\E`)
//fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Go."]

// 查找从行首开始，以空格结尾的字符串（非贪婪模式）
//reg = regexp.MustCompile(`(?U)^.* `)
//fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello "]

// 查找以空格开头，到行尾结束，中间不包含空格字符串
//reg = regexp.MustCompile(` [^ ]*$`)
//fmt.Printf("%q\n", reg.FindAllString(text, -1))
// [" Go."]

// 查找“单词边界”之间的字符串
//reg = regexp.MustCompile(`(?U)\b.+\b`)
//fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello" " 世界！" "123" " " "Go"]

// 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
//reg = regexp.MustCompile(`[^ ]{1,4}o`)
//fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello" "Go"]

// 查找 Hello 或 Go
//reg = regexp.MustCompile(`(?:Hell|G)o`)
//fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["Hello" "Go"]

// 查找 Hello 或 Go，替换为 Hellooo、Gooo
//reg = regexp.MustCompile(`(?PHell|G)o`)
//fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo"))
// "Hellooo 世界！123 Gooo."

// 交换 Hello 和 Go
//reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
//fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
// "Go 世界！123 Hello."

// 特殊字符的查找
//reg = regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]|]`)
//fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\123\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))
// "----------------------"
