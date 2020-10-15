package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8266")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}

		//create goroutine for each connect
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("client data: \n%v\n", str)
		newmessage := "ok\r\n"
		// send new string back to client
		fmt.Printf("server data: \n%v\n", newmessage)
		conn.Write([]byte(newmessage + "\n"))
	}
}
