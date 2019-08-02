package main

import (
	"bufio"
	"fmt"
	"github.com/googlebabaio/extsort/zujian"
	"os"
)

func main1() {

	p := zujian.ArraySource(3, 2, 5, 7, 6, 8, 1, 9)

	for {

		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}

func main2() {
	p := zujian.InMemSort(zujian.ArraySource(3, 2, 5, 7, 6, 8, 1, 9))
	for v := range p {
		fmt.Println(v)
	}
}

func main3() {
	p := zujian.Merge(zujian.InMemSort(zujian.ArraySource(3, 2, 5, 7)), zujian.InMemSort(zujian.ArraySource(6, 8, 1, 9)))
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

	p := zujian.RandomSource(n)
	zujian.WirteSink(bufio.NewWriter(file), p)

	fmt.Println("output successful!")

	file, err = os.Open(FILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = zujian.ReadSource(bufio.NewReader(file),-1)

	for v := range p {
		fmt.Println(v)
	}
}
