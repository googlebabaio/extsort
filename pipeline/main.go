package main

import (
	"encoding/binary"
	"fmt"
	"googlebaba.io/golangcookbook/zujian"
	"io"
	"math/rand"
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

func main() {
	p := zujian.Merge(zujian.InMemSort(zujian.ArraySource(3, 2, 5, 7)), zujian.InMemSort(zujian.ArraySource(6, 8, 1, 9)))
	for v := range p {
		fmt.Println(v)
	}
}

func ReadSource(reader io.Reader) <-chan int {

	out := make(chan int)

	go func() {
		buffer := make([]byte, 8)
		for {
			n, err := reader.Read(buffer)

			if err != nil {
				break
			}

			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
		}
		close(out)
	}()

	return nil
}

func WirteS(write io.WriteCloser, in <-chan int) {

	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		write.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
	}()
	return out
}
