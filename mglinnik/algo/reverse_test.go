package mglinnik

import "testing"

const failureString = "Expected %s to equal %s, but got %s"


func TestReverse(t *testing.T) {
	input := []string{
		"привет",
		"пустите в отпуск",
		"умоляю",
	}

	expected := []string{
		"тевирп",
		"ксупто в етитсуп",
		"юялому",
	}

	for i, x := range input {
		if Reverse(x) != expected[i] {
			t.Errorf(failureString, input[i], expected[i], Reverse(x))
		}
	}
}
