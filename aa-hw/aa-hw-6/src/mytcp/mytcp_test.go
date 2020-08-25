package mytcp

import (
	"testing"
)

func TestMyTCP(t *testing.T) {

	type parceData struct {
		tested   string
		expected string
	}

	p := []parceData{}

	p = append(p, parceData{
		`{"cmd":"foo", "args":[1,4,6,2,18,0]}`,
		`{"data":{"foo":[1,4,6,2,18,0]},"error":""}`})
	p = append(p, parceData{
		`{"cmd":"serv", "args":[1,4,6,2,18,0]}`,
		`{"data":{},"error":"wrong command"}`})
	p = append(p, parceData{
		`{"cmd":"foo"}`,
		`{"data":{"foo":null},"error":""}`})
	p = append(p, parceData{
		`{"cmd":"foo", "args":{Hello}}`,
		`{"data":null,"error":"invalid character 'H' looking for beginning of object key string"}`})

	for i, data := range p {
		r := request{}
		resp := r.parce(data.tested)
		result := string(resp.prepare())

		if result != data.expected {
			t.Errorf("Test MyTCP Error: i: %d, want: %s, get: %v\n", i, data.expected, result)
		}
	}
}
