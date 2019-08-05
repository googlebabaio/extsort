package main

import (
	"bufio"
	"fmt"
	"github.com/googlebabaio/extsort/pipeline/tools"
	"os"
)

func main1() {

	p := tools.ArraySource(3, 2, 5, 7, 6, 8, 1, 9)

	for {

		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}

func main2() {
	p := tools.InMemSort(tools.ArraySource(3, 2, 5, 7, 6, 8, 1, 9))
	for v := range p {
		fmt.Println(v)
	}
}

func main3() {
	p := tools.Merge(tools.InMemSort(tools.ArraySource(3, 2, 5, 7)), tools.InMemSort(tools.ArraySource(6, 8, 1, 9)))
	for v := range p {
		fmt.Println(v)
	}
}

func main() {

	const FILENAME  = "big.in"
	const n = 64
	file, err := os.Create(FILENAME)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	p := tools.RandomSource(n)

	writer := bufio.NewWriter(file)
	tools.WriteSink(writer, p)
	writer.Flush() //这一步很重要哦，不然数据不会被写入到文件中去


	fmt.Println("output successful!")

	//测试，从写入的文件中读取数据
	file, err = os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = tools.ReadSource(bufio.NewReader(file),-1)

	//输出前100行
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count>100{
			break
		}
	}
}
