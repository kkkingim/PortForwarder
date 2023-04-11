package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type ConnStr struct {
	LocalIP    string
	LocalPort  int
	RemoteIP   string
	RemotePort int
}

var connStr *ConnStr

func parseArgs() *ConnStr {
	args := os.Args
	cs := &ConnStr{}

	if len(args) <= 1 {
		return nil
	}

	a := args[1]
	as := strings.Split(a, ":")

	if len(as) != 4 {
		return nil
	}

	localPort, err := strconv.Atoi(as[1])
	if err != nil {
		return nil
	}
	remotePort, err := strconv.Atoi(as[3])
	if err != nil {
		return nil
	}

	localIP := as[0]
	remoteIP := as[2]

	cs.LocalIP = localIP
	cs.LocalPort = localPort
	cs.RemoteIP = remoteIP
	cs.RemotePort = remotePort

	return cs
}

func handleConn(conn net.Conn) {
	s := fmt.Sprintf("%s:%d", connStr.RemoteIP, connStr.RemotePort)
	conn2, err := net.Dial("tcp", s)
	if err != nil {
		log.Println("Error Dial to " + s)
		return
	}
	go func() {
		if _, err := io.Copy(conn2, conn); err != nil {
			return
		}
	}()
	if _, err := io.Copy(conn, conn2); err != nil {
		return
	}
	log.Println("Disconnected.")

}

func main() {
	connStr = parseArgs()
	if connStr == nil {
		fmt.Println("Usage: " + os.Args[0] + " LocalIP:LocalPort:RemoteIP:RemotePort")
		os.Exit(0)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", connStr.LocalIP, connStr.LocalPort))
	if err != nil {
		//listen panic
		log.Panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			go handleConn(conn)
		}
	}
}
