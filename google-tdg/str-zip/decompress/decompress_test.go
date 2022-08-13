package decompress

import "testing"

type TestCase struct {
	src    string
	result string
}

func TestDecompress(t *testing.T) {

	testCases := []TestCase{
		{
			src:    "abc",
			result: "abc",
		},
		{
			src:    "10[a]",
			result: "aaaaaaaaaa",
		},
		{
			src:    "2[3[a]b]",
			result: "aaabaaab",
		},
		{
			src:    "10[a]b",
			result: "aaaaaaaaaab",
		},
		{
			src:    "a",
			result: "a",
		},
		{
			src:    "",
			result: "",
		},
		{
			src:    "3[abc]4[ab]c",
			result: "abcabcabcababababc",
		},
	}

	for i := range testCases {
		r := Decompress_naive(testCases[i].src)
		if r != testCases[i].result {
			t.Errorf("error on test naive %v, expected %s but was %s", i,
				testCases[i].result, r)
		}
	}

	for i := range testCases {
		r := Decompress(testCases[i].src)
		if r != testCases[i].result {
			t.Errorf("error on test %v, expected %s but was %s", i,
				testCases[i].result, r)
		}
	}

}
