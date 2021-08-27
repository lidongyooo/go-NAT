package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
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
	addr, err := net.ResolveTCPAddr("tcp", serverIP + ":" + strconv.Itoa(sereverPort))
	serConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("连接到:%v, 等待转发请求... \n", serverIP + ":" + strconv.Itoa(sereverPort))
	serConn.SetDeadline(time.Now().Add(time.Hour))
	serConn.SetKeepAlive(true)

	locConn, err := net.Dial("tcp", ":" + strconv.Itoa(localPort))
	if err != nil {
		panic(err)
	}
	fmt.Printf("连接到:%v, 等待请求... \n", "127.0.0.1:" + strconv.Itoa(localPort))
	locConn.SetDeadline(time.Now().Add(time.Hour))

	for {
		request := make([]byte, 10240)
		n, err := serConn.Read(request)
		fmt.Println("===============start====================")
		fmt.Println(string(request))
		fmt.Println("==================end====================")
		//fmt.Println("serConn.read: ", err)

		go func() {
			_, err = locConn.Write(request[:n])
			//fmt.Println("localConn.write: ", err)

			response := make([]byte, 10240)
			n, err = locConn.Read(response)
			fmt.Println(string(response))
			fmt.Println("localConn.read: ", err)

			_, err = serConn.Write(response[:n])
			fmt.Println("serConn.write: ", err)
		}()
		
	}

}
