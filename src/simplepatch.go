package src

import (
	"fmt"
	"regexp"
)

func re() {
	//正则表达式
	const a = "2w niuwwwwww23mske@123.ded"
	re := regexp.MustCompile(`[a-zA-Z0-9]+@(.+\..+)`) //返回正则表达式, .代表任意字符，+代表一个或多个，*代表零个或多个
	//加（）表示提取
	match := re.FindString(a) //返回匹配子串
	match1 := re.FindAllStringSubmatch(a, -1)
	fmt.Println(match)
	fmt.Println(match1)
}
