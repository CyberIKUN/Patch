package src

import (
	"regexp"
)

type City struct {
	Link string
	Name string
}

func CityParser(buffer string, citylist *[]City) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\S*)" data-v-[a-z0-9^>^<]+>([\p{Han}]+)</a>`)
	match := re.FindAllStringSubmatch(buffer, -1)
	for _, value := range match {
		city := City{Link: value[1],
			Name: value[2]}
		*citylist = append(*citylist, city)
	}
}
