package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"patch/src"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	cityList := make([]int, 400)
	for i := 0; i < len(cityList); i++ {
		wait.Add(1)
		go func(a int) {
			for j := 2; ; {
				fmt.Println(a)
				if j == 399 {
					break
				}
				j++
			}
			wait.Done()
		}(i)
	}
	wait.Wait()
}

func test12() {
	www := make(chan int) //无缓冲通道，有一个读一个
	go func() {
		for i := 0; i < 399; i++ {
			www <- 1 //程序阻塞在这
			fmt.Println(i)
		}
	}()
	fmt.Println("--------------------------------------------")
	fmt.Println(cap(www))
	for i := 0; i < 399; i++ {
		<-www
	}
}

/*简单读取*/
func test5() {
	file, err := os.Open("D:\\Users\\lixuyong\\GolandProjects\\patch\\src\\test\\a.txt")
	//os.Open("./a.txt")不能打开不存在的文件，打开存在的文件，默认权限为只读
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 10000)
	count, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println("读取的字节数" + strconv.Itoa(count))
	fmt.Println(string(buffer))

	file.Close()
}

/*1.高效读取*/
func test6() {
	file, err := os.Open("D:\\Users\\lixuyong\\GolandProjects\\patch\\src\\test\\a.txt")
	//os.Open("./a.txt")不能打开不存在的文件，打开存在的文件，权限为只读
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //读取到换行符停止，读取出来的字符串包括换行符
		line = strings.ReplaceAll(line, "\n", "")
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(line)
	}

	file.Close()
}

/*2.高效读取*/
func test11() {
	file, err := os.Open("D:\\Users\\lixuyong\\GolandProjects\\patch\\src\\test\\a.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*一次性读取，小文件适合*/
func test7() {
	//不用自己打开文件和关闭文件
	buffer, err := ioutil.ReadFile("D:\\Users\\lixuyong\\GolandProjects\\patch\\src\\test\\a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer))
}

/*高效写文件*/
func test8() {
	//涉及到写文件，一定要用os.OpenFile()
	file, err := os.OpenFile(`D:\Users\lixuyong\GolandProjects\patch\src\test\a.txt`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0766)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	writeCount, err := writer.WriteString("abc")
	if err != nil {
		panic(err)
	}

	writer.Flush()

	fmt.Println(writeCount)
}

/*简单将一个文件的内容写入另一个文件*/
func test9() {
	content, err := ioutil.ReadFile("D:\\Users\\lixuyong\\GolandProjects\\patch\\src\\test\\a.txt")
	if err != nil {
		panic(err)
	}
	//不用自己打开和关闭文件，覆盖写，不存在就创建
	err = ioutil.WriteFile("D:\\Users\\lixuyong\\GolandProjects\\patch\\src\\test\\b.txt", content, 0766)
	if err != nil {
		panic(err)
	}
}

/*判断文件是否存在*/
func test10() {
	_, err := os.Stat("d:\\aaa.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("文件不存在")
	}
}
func test3() {
	result := src.Engine("https://www.zhenai.com/zhenghun/akesu/1")
	var personlist []src.Person
	src.PersonParser(&personlist, result)
	fmt.Println(personlist)
}

func test2() {
	var result = src.Engine("https://www.zhenai.com/zhenghun/akesu/1")
	fmt.Println(result)
}
func test1() {
	re := regexp.MustCompile(`[\d|\d-\d]+`)
	match := re.FindAllString("5000,5000-8000", -1)
	fmt.Println(match)
}
func test4() {
	match := regexp.MustCompile(`([\p{Han}a-zA-Z0-9]+)`)
	result := match.FindAllStringSubmatch("心藏", -1)
	fmt.Println(result)
}
