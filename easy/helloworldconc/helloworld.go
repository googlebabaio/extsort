package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go printHello(i, ch)
	}

	for {
		msg := <-ch
		println(msg)
	}
}

func printHello(n int, ch chan string) {
	for {
		ch <- fmt.Sprintf("哈哈哈哈:%d", n)
	}

}
