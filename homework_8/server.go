package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Command struct {
	Cmd  string `json:"cmd"`
	Args []int  `json:"args"`
}

func fromJson(lineToUnmarshal string) (command Command){
	err := json.Unmarshal([]byte(lineToUnmarshal), &command)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return command
}
func main() {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", message)
		command := fromJson(message)
		if command.Cmd != "foo" {
			conn.Write([]byte("Wrong command in json !" + "\n"))
		} else {
			conn.Write([]byte(strings.Trim(strings.Join(strings.Split(fmt.Sprint(command.Args), " "), ", " ), "[]") + "\n"))
		}

	}
}