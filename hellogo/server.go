package main
import (
	"fmt"
	"net"
	 "log"
	"os"
)

func main() {
	//建立socket，监听端口
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()

	log.Printf("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		log.Printf(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}
//处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		log.Printf(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}