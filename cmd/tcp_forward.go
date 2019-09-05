package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	var srcAddress string
	var srcPort, destPort int
	flag.StringVar(&srcAddress, "s", "", "server address")
	flag.IntVar(&srcPort, "sp", 0, "server port")
	flag.IntVar(&destPort, "dp", 0, "local port")
	ln, err := net.Listen("tcp",  ":" + destPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("start forwarding tcp from %s:%d to :%d ...\n", srcAddress, srcPort, destPort)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	closed := false
	defer func() {
		if !closed {
			conn.Close()
		}
	}()

}
