package aahw4

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func (root *Node) print(w io.Writer, ns int, ch rune) {
	if root == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, root.data)
	root.left.print(w, ns-2, 'L')
	root.right.print(w, ns+2, 'R')
}

func TestHeight(t *testing.T) {

	type testValue struct {
		array []int
		want  int
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{
		[]int{3, 5, 2, 1, 4, 6, 7},
		3,
	})

	testValueArr = append(testValueArr, testValue{
		[]int{15},
		0,
	})

	testValueArr = append(testValueArr, testValue{
		[]int{2, 1, 7, 5, 4},
		3,
	})

	var tree *Tree
	for i, tVal := range testValueArr {

		tree = &Tree{}
		for _, data := range tVal.array {
			tree = tree.insert(data)
		}
		fmt.Println("##############################")
		tree.root.print(os.Stdout, 15, 'M')
		result := tree.root.height()
		if tVal.want != result {
			t.Errorf("Error Test height. Iteration = %d, Want = %d and Got = %v.\n", i, tVal.want, result)
		} else {
			fmt.Printf("Ok Test height. Iteration = %d, Want = %d and Got = %v.\n", i, tVal.want, result)
		}
	}

}

func TestLCA(t *testing.T) {

	type testValue struct {
		array []int
		v1    int
		v2    int
		want  int
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{
		[]int{4, 2, 3, 1, 7, 6},
		1, 7,
		4,
	})

	testValueArr = append(testValueArr, testValue{
		[]int{1, 2},
		1, 2,
		1,
	})

	testValueArr = append(testValueArr, testValue{
		[]int{5, 3, 8, 2, 4, 6, 7},
		7, 3,
		5,
	})

	var tree *Tree
	for i, tVal := range testValueArr {

		tree = &Tree{}
		for _, data := range tVal.array {
			tree = tree.insert(data)
		}
		fmt.Println("##############################")
		tree.root.print(os.Stdout, 15, 'M')
		result := tree.root.lca(tVal.v1, tVal.v2)
		if tVal.want != result.data {
			t.Errorf("Error Test lca. Iteration = %d, Want = %d and Got = %v.\n", i, tVal.want, result.data)
		} else {
			fmt.Printf("Ok Test lca. Iteration = %d, Want = %d and Got = %v.\n", i, tVal.want, result.data)
		}
	}
}

func TestBST(t *testing.T) {

	type testValue struct {
		array []int
		want  bool
	}

	var testValueArr []testValue

	testValueArr = append(testValueArr, testValue{
		[]int{1, 2, 3, 4, 5, 6, 7},
		true,
	})

	testValueArr = append(testValueArr, testValue{
		[]int{1, 2, 4, 3, 5, 6, 7},
		false,
	})

	testValueArr = append(testValueArr, testValue{
		[]int{3, 5, 7, 9, 11, 13, 15},
		true,
	})

	var tree *Tree
	for i, tVal := range testValueArr {

		tree = &Tree{}
		for _, data := range tVal.array {
			tree = tree.insert(data)
		}
		fmt.Println("##############################")
		tree.root.print(os.Stdout, 15, 'M')
		result := tree.root.bst()
		if tVal.want != result {
			t.Errorf("Error Test bst. Iteration = %d, Want = %v and Got = %v.\n", i, tVal.want, result)
		} else {
			fmt.Printf("Ok Test bst. Iteration = %d, Want = %v and Got = %v.\n", i, tVal.want, result)
		}
	}
}
