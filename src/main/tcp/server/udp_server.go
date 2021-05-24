package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running...")
	conn, err := net.ListenPacket("udp", ":10000")
	if err != nil {
		fmt.Errorf("failed to create soket")
		panic(err)
	}
	// 関数終了時にコネクションを閉じる
	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received from %v: %v\n",
			remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"),
			remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
