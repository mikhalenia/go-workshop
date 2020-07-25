package frogstair

import "testing"

func TestReverse(t *testing.T) {
	data := []string{
		"The quick brown 狐 jumped over the lazy 犬",
		"Привет 世界!",
		"Съешь ещё этих мягких французских булок, да выпей же чаю",
	}

	expected := []string{
		"犬 yzal eht revo depmuj 狐 nworb kciuq ehT",
		"!界世 тевирП",
		"юач еж йепыв ад ,колуб хиксзуцнарф хикгям хитэ ёще ьшеъС",
	}

	for i, d := range data {
		if expected[i] != Reverse(d) {
			t.Errorf("ERROR: Want: %s, got %s", expected[i], Reverse(d))
		}
	}

}
