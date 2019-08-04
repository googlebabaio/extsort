package tools

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sort"
	"time"
)

var startTime time.Time

func InitTime()  {
	startTime=time.Now()
}

func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		//读取到内存中
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		//内存中排序
		fmt.Println("Read done",time.Now().Sub(startTime))
		sort.Ints(a)
		fmt.Println("InMemSort done",time.Now().Sub(startTime))

		//将排好序的值，再扔回给 channel
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

/**
合并两个输入并输出
*/
func Merge(in1, in2 <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		v1, ok1 := <-in1
		v2, ok2 := <-in2

		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done",time.Now().Sub(startTime))
	}()

	return out
}

func ReadSource(reader io.Reader, chunkSize int) <-chan int {

	out := make(chan int)

	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n

			//注意和下面的判断的顺序
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}

			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()

	return out
}

func WriteSink(write io.Writer, in <-chan int) {

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
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {

	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2

	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))

}

func NetworkSink(addr string , in <-chan int){
	listener , err := net.Listen("tcp",addr)
	if err != nil{
		panic(err)
	}
	go func(){
		defer listener.Close()
		conn , err := listener.Accept()
		if err != nil{
			panic(err)
		}
		defer conn.Close()
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriteSink(writer,in)
	}()

}

func NetworkSource(addr string) <-chan int{
	out := make(chan int)
	go func(){
		conn , err := net.Dial("tcp",addr)
		if err != nil{
			panic(err)
		}
		r := ReadSource(bufio.NewReader(conn),-1)
		for v := range r{
			out <- v
		}
		close(out)
	}()
	return out
}

