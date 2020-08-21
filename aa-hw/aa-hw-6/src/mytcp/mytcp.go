package mytcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type request struct {
	Cmd  string `json:"cmd"`
	Args []int  `json:"args"`
}

type response struct {
	Data  map[string][]int `json:"data"`
	Error string           `json:"error"`
}

func (r *request) unmarshal(raw string) (err error) {
	err = json.Unmarshal([]byte(raw), &r)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (r *request) parce(raw string) (resp response) {
	data := map[string][]int{}

	if err := r.unmarshal(raw); err != nil {
		resp.Error = err.Error()
		return
	}

	switch r.Cmd {
	case "foo":
		data[r.Cmd] = r.Args
	default:
		resp.Error = "wrong command"
	}
	resp.Data = data
	return
}

func (r *response) prepare() []byte {
	w, _ := json.Marshal(r)
	return w
}

func Listen(port string) {
	var err error
	var ln net.Listener
	var conn net.Conn

	if ln, err = net.Listen("tcp", ":"+port); err != nil {
		panic(err)
	}
	fmt.Printf("Listen port %s\n", port)

	if conn, err = ln.Accept(); err != nil {
		panic(err)
	}

	defer func() {
		conn.Close()
		ln.Close()
	}()

	for {
		var err error
		var raw string
		resp := response{}

		if raw, err = bufio.NewReader(conn).ReadString('\n'); err == nil {
			r := request{}
			resp = r.parce(raw)
		}

		if err != nil {
			resp.Error = err.Error()
		}
		preparedResp := resp.prepare()

		fmt.Printf("response: %v\n", string(preparedResp))
		conn.Write(preparedResp)
	}
}
