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

	const FILENAME  = "small.in"
	const n = 50
	file, err := os.Create(FILENAME)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	p := tools.RandomSource(n)
	tools.WirteSink(bufio.NewWriter(file), p)

	fmt.Println("output successful!")

	file, err = os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = tools.ReadSource(bufio.NewReader(file),-1)

	for v := range p {
		fmt.Println(v)
	}
}
