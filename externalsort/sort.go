package main

import (
	"bufio"
	"fmt"
	"github.com/googlebabaio/extsort/pipeline/tools"
	"os"
	"strconv"
)

func main() {
	infile := "small.in"
	outfile := "small.out"
	p := createPipeline(infile, 512, 4)
	writeToFile(p, outfile)
	printFile(outfile)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := tools.ReadSource(file, -1)

	for v := range p {
		fmt.Println(v)
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	tools.WriteSink(writer, p)
}

//chunkCount 桶的数量
//chunkSize 桶的大小，用 filesize 除以 chunkcount
//fileSize 文件的大小
func createPipeline(filename string, fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := tools.ReadSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, tools.InMemSort(source))

	}

	return tools.MergeN(sortResults...)
}

func createNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount

	tools.InitTime()
	sortAddr := []string{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := tools.ReadSource(
			bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000+i)
		// 塞给网络服务器
		tools.NetworkSink(addr, tools.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	// 从网络服务器取
	sortResults := [] <-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults,
			tools.NetworkSource(addr))
	}
	return tools.MergeN(sortResults...)
}
