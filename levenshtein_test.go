package levenshtein

import (
	"fmt"
	"testing"
)

var distanceTests = []struct {
	str_1    string
	str_2    string
	expected int
}{
	{"កកាត", "កកាត", 0},
	{"កកោស", "ក", 3},
	{"ឯថា", "ឯក", 3},
	{"ឯថា", "ឯក", 3},
	{"ឯកឯង", "ឯកកង", 2},
	{"hello", "", 5},
	{"hello", "hello", 0},
	{"ab", "aa", 2},
	{"ab", "aaa", 3},
	{"bbb", "a", 4},
	{"kitten", "sitting", 5},
	{"distance", "difference", 8},
	{"levenshtein", "frankenstein", 9},
	{"resume and cafe", "resumes and cafes", 2},
}

func TestDistance(t *testing.T) {

	for index, distanceTest := range distanceTests {
		result := Distance(distanceTest.str_1, distanceTest.str_2)
		if result != distanceTest.expected {
			output := fmt.Sprintf("%v \t ចំងាយរវាង %v និង %v គួរតែ %v ប៉ុន្តែ %v.",
				index, distanceTest.str_1, distanceTest.str_2, distanceTest.expected, result)
			t.Errorf(output)
		}
	}
}
