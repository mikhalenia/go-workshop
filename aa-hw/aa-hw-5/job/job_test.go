package job

import (
	"reflect"
	"testing"
)

func TestJobPush(t *testing.T) {

	testQueue := new(NewQueue)

	testPush := []struct {
		Content []string
		Push []string
		Final []string
	}{
		{
			[]string{"1", "2", "3", "4"},
			[]string {"5", "6", "7"},
			[]string{"1", "2", "3", "4", "5", "6", "7"},
		},
		{
			[]string{},
			[]string {"3","4", "5"},
			[]string{"3", "4","5"},
		},
	}



	for i, push := range testPush {
		testQueue.Content = push.Content
		for _, el := range push.Push {
			testQueue.Push(el)
		}
		if !reflect.DeepEqual(testQueue.Content, push.Final) {
			t.Errorf("error test case %s: want %v, got %v", i, push.Final, testQueue.Content)
		}
	}
}

func TestJobPop(t *testing.T) {

	testQueue := new(NewQueue)

	testPop := []struct {
		Content  []string
		Expected string
		Error    string
		Final    []string
	}{
		{
			[]string{"1", "2", "3", "4"},
			"1",
			"",
			[]string{"2", "3", "4"},
		},
		{
			[]string{},
			"",
			"empty queue",
			[]string{},
		},
	}

	for i, pop := range testPop {
		testQueue.Content = pop.Content
		j, err := testQueue.Pop()
		if pop.Error == "" && err != nil {
			t.Errorf("error case %s: unexpected error %v", i, err.Error())
		}
		if pop.Error != "" && err == nil {
			t.Errorf("error case %s: unexpected error %v", i, pop.Error)
		}
		if j != pop.Expected || !reflect.DeepEqual(testQueue.Content, pop.Final) {
			t.Errorf("error case %s: wanted %v, got %v, expected value from pop %v, got from pop %v",
				i,
				pop.Final,
				testQueue.Content,
				j,
				pop.Expected)
		}
	}
}