package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

var (
	localPort int
	sereverPort int
	serverIP string
)

func init() {
	flag.IntVar(&localPort, "l", 8080, "本地服务端口")
	flag.IntVar(&sereverPort, "s", 12001, "远程服务端口")
	flag.StringVar(&serverIP, "h", "127.0.0.1", "远程服务地址")
}

func main() {
	flag.Parse()

	serConn, err := net.Dial("tcp", serverIP + ":" + strconv.Itoa(sereverPort))
	if err != nil {
		panic(err)
	}
	fmt.Printf("连接到:%v, 等待转发请求... \n", serverIP + ":" + strconv.Itoa(sereverPort))

	locConn, err := net.Dial("tcp", ":" + strconv.Itoa(localPort))
	if err != nil {
		panic(err)
	}
	fmt.Printf("连接到:%v, 等待请求... \n", "127.0.0.1:" + strconv.Itoa(localPort))

	for {
		fmt.Println("start")
		request := make([]byte, 10240)
		n, _ := serConn.Read(request)
		locConn.Write(request[:n])

		response := make([]byte, 10240)
		n, _ = locConn.Read(response)
		serConn.Write(response[:n])
		fmt.Println("end")
	}

}
