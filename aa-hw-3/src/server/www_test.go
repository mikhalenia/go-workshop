package server

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestWWW(t *testing.T) {

	InitServer()
	go StartWebServer()

	for {
		err, str := runCommand("curl", "http://localhost")
		if err == nil && str != "" {
			break
		}
		time.Sleep(1 * time.Second)
	}

	testOrders(t)

}

func testOrders(t *testing.T) {
	for i := 1; i < 10; i++ {
		val := fmt.Sprintf("http://localhost/?createorder=0&auto=%d", i)
		err, _ := runCommand("curl", val)
		isError(t, err)
	}

	if len(orders.Table) <= 0 {
		t.Errorf("Error Test Len Orders Table. Want > 0 and Got = '%v'", orders.Table)
	}

	val := fmt.Sprintf("http://localhost/?paidorder=%d&type=_", 5)
	err, _ := runCommand("curl", val)
	isError(t, err)

}

func runCommand(cmdName string, arg ...string) (error, string) {
	cmd := exec.Command(cmdName, arg...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

	return err, outStr
}

func isError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error Test: %s.", err.Error())
	}
}
