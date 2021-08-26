package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
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
	serConn, err := serLis.Accept()

	for {
		fmt.Println("start")

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
		netConn.SetDeadline(time.Now().Add(time.Second * 10))
		n, _ := netConn.Read(request)
		serConn.Write(request[:n])

		response := make([]byte, 10240)
		serConn.SetDeadline(time.Now().Add(time.Second * 10))
		n, _ = serConn.Read(response)
		netConn.Write(response[:n])
		fmt.Println("end")
	}

}

func netHandler()  {

}

func serHandler() {

}
