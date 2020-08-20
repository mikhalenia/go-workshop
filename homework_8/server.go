package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
)
import "fmt"
import "bufio"
import "strings"
type Command struct {
	Cmd  string `json:"cmd"`
	Args []int  `json:"args"`
}
func toJson(data Command) string {
	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return string(result)
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

	// Устанавливаем прослушивание порта
	ln, _ := net.Listen("tcp", ":8081")

	// Открываем порт
	conn, _ := ln.Accept()

	// Запускаем цикл
	for {
		// Будем прослушивать все сообщения разделенные \n
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// Распечатываем полученое сообщение
		fmt.Print("Message Received:", message)
		command := fromJson(message)
		if command.Cmd != "foo" {
			conn.Write([]byte("Wrong command in json !" + "\n"))
		} else {
			conn.Write([]byte(strings.Trim(strings.Join(strings.Split(fmt.Sprint(command.Args), " "), ", " ), "[]") + "\n"))
		}

	}
}