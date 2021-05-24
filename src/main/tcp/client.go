package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp4", ":10000")
	if err != nil {
		fmt.Errorf("failed to create connection")
		panic(err)
	}
	// 関数終了時にコネクションを閉じる
	defer conn.Close()
	fmt.Println("Send")
	_, err = conn.Write([]byte("Hello World"))
	if err != nil {
		fmt.Errorf("failed to send server")
		panic(err)
	}

	fmt.Println("Read")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s\n", string(buffer[:length]))
}
