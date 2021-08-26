package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

var (
	netPort int
	serverPort int
)

func init() {
	flag.IntVar(&netPort, "n", 12000, "面向外网端口")
	flag.IntVar(&serverPort, "s", 12001, "客户端连接端口")
}

func main() {
	flag.Parse()

	netLis, err := net.Listen("tcp", ":" + strconv.Itoa(netPort))
	if err != nil {
		panic(err)
	}
	fmt.Printf("监听:%d端口, 等待网络请求... \n", netPort)

	serLis, err := net.Listen("tcp", ":" + strconv.Itoa(serverPort))
	if err != nil {
		panic(err)
	}
	fmt.Printf("监听:%d端口, 等待客户端连接... \n", serverPort)

	for {
		serConn, err := serLis.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		netConn, err := netLis.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		request := make([]byte, 10240)
		netConn.Read(request)
		serConn.Write(request)

		response := make([]byte, 10240)
		serConn.Read(response)
		netConn.Write(response)
	}

}

func netHandler()  {

}

func serHandler() {

}
