package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

// Start starts the server
func Start(port string) {
	fmt.Println("Server started")

	l, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("Accepted connection")

	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		var cmd map[string]interface{}
		err = json.Unmarshal([]byte(data), &cmd)
		if err != nil {
			writeBack(c, "invalid JSON", "")
			continue
		}

		if cmd["cmd"] == "exit" {
			writeBack(c, "", "bye")
			fmt.Println("Stopped server")
			c.Close()
			l.Close()
			break
		}

		res, err := parseCommand(
			cmd["cmd"].(string),
			cmd["args"],
		)
		if err != nil {
			writeBack(c, err.Error(), "")
			continue
		}

		writeBack(c, "", res)
	}
}

func writeBack(w net.Conn, error, data string) {
	var response struct {
		Error string `json:"error"`
		Data  string `json:"data"`
	}
	response.Error = error
	response.Data = data
	d, _ := json.Marshal(response)
	w.Write(d)
	w.Write([]byte{10})
}
