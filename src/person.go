package src

import (
	"regexp"
	"strconv"
	"sync"
)

type Person struct {
	Name          string
	Sex           string
	Address       string
	Salary        string
	Age           int
	Height        int
	MaritalStatus string
}

var mutex sync.RWMutex

func PersonParser(personlist *[]Person, result string) {
	mailre := regexp.MustCompile(`<tbody><tr><th><a href="http://album.zhenai.com/u/[\d]+" target="_blank">([\p{Han}a-zA-Z0-9]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([\p{Han}]{2})</td> <td><span class="grayL">居住地：</span>([\p{Han}]+)</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>([\d]{2})</td>  <td><span class="grayL">月   薪：</span>([\d|\d-\d]+[\p{Han}]+)</td></tr> <tr><td width="180"><span class="grayL">婚况：</span>([\p{Han}]{2})</td> <td width="180"><span class="grayL">身   高：</span>([\d]+)</td></tr></tbody>`)
	mailmatch := mailre.FindAllStringSubmatch(result, -1)
	for _, value := range mailmatch {
		age, err := strconv.Atoi(value[4])
		if err != nil {
			panic(err)
			return
		}
		height, err := strconv.Atoi(value[7])
		if err != nil {
			panic(err)
			return
		}
		person := Person{Name: value[1],
			Sex:           value[2],
			Address:       value[3],
			Age:           age,
			Salary:        value[5],
			MaritalStatus: value[6],
			Height:        height,
		}
		mutex.Lock()
		*personlist = append(*personlist, person)
		mutex.Unlock()
	}
}
