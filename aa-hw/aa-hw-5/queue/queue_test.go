package queue

import (
	"testing"
	"reflect"
)
func TestQueuePush(t *testing.T) {

	testQueue := new(Queue)

	testPush := []struct {
		Content []int
		Push []int
		Final []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int {5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{},
			[]int {3, 4, 5},
			[]int{3, 4, 5},
		},
	}



	for i, push := range testPush {
		testQueue.Content = push.Content
		for _, el := range push.Push {
			testQueue.Push(el)
		}
		if !reflect.DeepEqual(testQueue.Content, push.Final) {
			t.Errorf("error test case %d: want %v, got %v", i, push.Final, testQueue.Content)
		}
	}
}

func TestQueuePop(t *testing.T) {

	testQueue := new(Queue)

	testPop := []struct {
		Content []int
		Expected int
		Error string
		Final []int
	}{
		{
			[]int{1, 2, 3, 4},
			1,
			"",
			[]int{2, 3, 4},
		},
		{
			[]int{},
			0,
			"empty queue",
			[]int{},
		},
	}

	for i, pop := range testPop {
		testQueue.Content = pop.Content
		j, err := testQueue.Pop()
		if pop.Error == "" && err != nil{
			t.Errorf("error case %d: unexpected error %v",i, err.Error())
		}
		if pop.Error != "" && err == nil {
			t.Errorf("error case %d: unexpected error %v",i, pop.Error)
		}
		if j != pop.Expected || !reflect.DeepEqual(testQueue.Content, pop.Final) {
			t.Errorf("error case %d: wanted %v, got %v, expected value from pop %v, got from pop %v",
				i,
				pop.Final,
				testQueue.Content,
				j,
				pop.Expected)
		}
	}
}