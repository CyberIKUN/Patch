package main

import (
	"bufio"
	"fmt"
	"os"
	"patch/src"
	"strconv"
	"sync"
)

var (
	rootURL    = "https://www.zhenai.com/zhenghun"
	cityList   []src.City
	personList []src.Person
	wait       sync.WaitGroup
)

func main() {
	urls := src.Engine(rootURL)
	src.CityParser(urls, &cityList)
	for i := 0; i < len(cityList); i++ {
		response := src.Engine(cityList[i].Link)
		src.PersonParser(&personList, response)
		wait.Add(1)
		go func(a int) {
			for j := 2; ; {
				url := fmt.Sprintf(cityList[a].Link+"/%d", j)
				res := src.Engine(url)
				if res != "" {
					src.PersonParser(&personList, res)
					j++
					fmt.Println("正在加入人员")
				} else {
					break
				}
			}
			wait.Done()
		}(i)
	}

	wait.Wait()
	//写入文本
	file, err := os.OpenFile("./wuhu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < len(personList); i++ {
		count, err := writer.WriteString(personList[i].Name + "," + personList[i].Sex + "," +
			personList[i].Address + "," + personList[i].Salary + "," +
			strconv.Itoa(personList[i].Age) + "," +
			strconv.Itoa(personList[i].Height) +
			personList[i].MaritalStatus)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(count)
	}
	writer.Flush()

}
