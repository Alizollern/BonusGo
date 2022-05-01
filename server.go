package main

import (
	"context"
	"net"
	"fmt"
	"bufio"
	"strings"
	"log"
)

func handleConnection(net.Conn) {}

func run(ctx context.Context) error {

	var lc net.ListenConfig

	ln, err := lc.Listen(ctx, "tcp", ":4041")
	if err != nil {
		return err
	}
	fmt.Println("Launching server...")
	conn, _ := ln.Accept()
	for {
    	message, _ := bufio.NewReader(conn).ReadString('\n')
    	fmt.Print("Message Received:", string(message))
    	str := "Hello "+string(message)+"\n"
    	conn.Write([]byte(str))
    	newmessage := strings.ToUpper(message)
    	conn.Write([]byte(newmessage + "\n"))
  	}
  	return nil
}

func main(){
	ctx, _ := context.WithCancel(context.Background())
	err := run(ctx)
	if err == context.Canceled {
		log.Printf("Shutdown is done")
	} else if err != nil {
		log.Printf("ERROR: %v", err)
	}
}

