package main

import (
	"reflect"
	"testing"
)

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

func TestShuffle(t *testing.T) {
	problems := []Problem{
		{question: "qa", answer: "aa"},
		{question: "qb", answer: "ab"},
		{question: "qc", answer: "ac"},
		{question: "qd", answer: "ad"},
		{question: "qe", answer: "ae"},
		{question: "qf", answer: "af"},
		{question: "qg", answer: "ag"},
	}
	notWant := make([]Problem, len(problems))
	copy(notWant, problems)

	Shuffle(problems)
	if reflect.DeepEqual(problems, notWant) {
		t.Errorf("Expeced problems %s not to be same as shuffled problems %s", problems, notWant)
	}
}
