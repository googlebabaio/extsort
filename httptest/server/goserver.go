package server

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var quitSemaphore chan bool
var count = 0
var gLocker sync.Mutex //全局锁

func main() {
	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "0.0.0.0:6666")

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		gLocker.Lock()
		count++
		gLocker.Unlock()
		fmt.Printf("连接数：%d , %s \n ", count, tcpConn.RemoteAddr().String())
		//fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn) {
	//ipStr := conn.RemoteAddr().String()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Read error:%v\n", err)
			gLocker.Lock()
			count--
			gLocker.Unlock()
			<-quitSemaphore
		}

		fmt.Println(string(message))

		// msg := time.Now().Format("2006-01-02 15:04:05") + "\n"
		// b := []byte(msg)
		// conn.Write(b)
	}
}
