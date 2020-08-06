package huffman

import (
	"fmt"
	//	"io"
	//	"os"
	"testing"
)

func (t *Tree) createHuffmanTree(array string, value string, code map[string]string) *Tree {
	//
}

func TestHuffman(t *testing.T) {

	type testValue struct {
		array string
		want  string
		code  map[string]string
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{
		"1001011",
		"ABACA",
		map[string]string{
			"1":  "A",
			"00": "B",
			"01": "C",
		},
	})

	testValueArr = append(testValueArr, testValue{
		"01111001100011010111100",
		"ABRACADABRA",
		map[string]string{
			"0":    "A",
			"111":  "B",
			"1100": "C",
			"1101": "D",
			"10":   "R",
		},
	})

	var tree *Tree
	for i, tVal := range testValueArr {
		tree = &Tree{}
		//for _, data := range tVal.array {
		tree = tree.createHuffmanTree(tVal.array, tVal.value, tVal.code)
		//}

		result := tree.root.huffman(tVal.array)
		if tVal.want != result {
			t.Errorf("Error Test huffman. Iteration = %d, Want = %v and Got = %v.\n", i, tVal.want, result)
		} else {
			fmt.Printf("Ok Test huffman. Iteration = %d, Want = %v and Got = %v.\n", i, tVal.want, result)
		}
	}
}
