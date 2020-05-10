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
}

func TestDistance(t *testing.T) {

	for index, distanceTest := range distanceTests {
		result := Distance(distanceTest.str_1, distanceTest.str_2)
		if result != distanceTest.expected {
			output := fmt.Sprintf("%v \t distance between %v and %v should be %v but was %v.",
				index, distanceTest.str_1, distanceTest.str_2, distanceTest.expected, result)
			t.Errorf(output)
		}
	}
}
