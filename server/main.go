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

	addr, err := net.ResolveTCPAddr("tcp", ":" + strconv.Itoa(netPort))
	netLis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("监听:%d端口, 等待网络请求... \n", netPort)

	addr, err = net.ResolveTCPAddr("tcp", ":" + strconv.Itoa(serverPort))
	serLis, err := net.ListenTCP("tcp", addr)
	fmt.Printf("监听:%d端口, 等待客户端连接... \n", serverPort)
	serConn, err := serLis.AcceptTCP()
	serConn.SetDeadline(time.Now().Add(time.Hour))
	serConn.SetKeepAlive(true)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {

		netConn, err := netLis.AcceptTCP()
		netConn.SetKeepAlive(true)
		netConn.SetDeadline(time.Now().Add(time.Hour))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		go func() {
			for {
				request := make([]byte, 10240)
				n, err := netConn.Read(request)
				fmt.Println("netConn.read", err)
				_, err = serConn.Write(request[:n])
				fmt.Println("serConn.write", err)

				response := make([]byte, 10240)
				n, err = serConn.Read(response)
				fmt.Println("serConn.read", err)

				_, err = netConn.Write(response[:n])
				fmt.Println("netConn.write", err)

			}
		}()
	}

}
