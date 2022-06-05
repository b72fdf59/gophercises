package main

import "testing"

func TestParseAnswer(t *testing.T) {
	cases := []struct {
		testString string
		want       string
	}{
		{
			testString: " Hello this is new ",
			want:       "Hello this is new",
		},
	}

	for _, test := range cases {
		got := ParseAnswer(test.testString)
		if test.want != got {
			t.Errorf("Expeced %s, but got %s", test.want, got)
		}

	}
}
