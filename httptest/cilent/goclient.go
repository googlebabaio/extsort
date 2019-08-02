package cilent

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphore chan bool

func main() {

	for a := 0; a < 60000; a++ {
		var tcpAddr *net.TCPAddr

		tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.4.108:6666")

		conn, error := net.DialTCP("tcp", nil, tcpAddr)
		if error != nil {
			fmt.Printf("error is: %v\n", error)
		} else {
			fmt.Printf("connected!: %d\n", a)
		}
		if conn != nil {
			go onMessageRecived(conn)
			msg := time.Now().Format("2006-01-02 15:04:05") + "\n"
			b := []byte(msg)
			conn.Write(b)
		}
	}

	<-quitSemaphore

}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, error := reader.ReadString('\n')
		if error != nil {
			fmt.Printf("error is: %v\n", error)
		} else {
			fmt.Println(msg)
		}

	}
}
