package hw6

import "testing"

func EncodeTest(t *testing.T) {
	testData := []struct {
		Test struct {
			A string
			B int
			C struct {
				D int
				E int
			}
		}
		Expected string
	}{
		{
			Test: struct {
				A string
				B int
				C struct {
					D int
					E int
				}
			}{
				A: "Hello",
				B: 10,
				C: struct {
					D int
					E int
				}{
					D: 15,
					E: 25,
				},
			},
			Expected: `{"A":"Hello","B":10,"C":{"D":15,"E":25}}`,
		},
		{
			Test: struct {
				A string
				B int
				C struct {
					D int
					E int
				}
			}{
				A: "World!",
				B: 49,
				C: struct {
					D int
					E int
				}{
					D: 135,
					E: 53,
				},
			},
			Expected: `{"A":"World!","B":49,"C":{"D":135,"E":53}}`,
		},
	}

	for _, el := range testData {
		test, err := Encode(el.Test)
		if err != nil {
			panic(err)
		}
		if test != el.Expected {
			t.Errorf("ERROR: want %s, got %s", el.Expected, test)
		}
	}

}

// ????

//func DecodeTest(t *testing.T) {
//	testData := []struct {
//		Input    string
//		Expected struct {
//			A int
//			B string
//			C map[string]interface{}
//		}
//	}{
//		{
//			Input: `{"A":15,"B":"Hello world!","C":{"D":"foo","E":"bar","F":"faz"}}`,
//			Expected: struct {
//				A int
//				B string
//				C map[string]interface{}
//			}{
//				A: 15,
//				B: "Hello world!",
//				C: map[string]interface{}{
//					"D": "foo",
//					"E": "bar",
//					"F": "baz",
//				},
//			},
//		},
//	}
//
//	for _, el := range testData {
//		v := struct {
//			A int
//			B string
//			C map[string]interface{}
//		}{}
//
//		v, err := Decode(el.Input, v)
//
//	}
//}
