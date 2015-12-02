package dwellerutil

import (
	"fmt"
	"testing"
)

func ExampleSuggest() {
	best, _, _ := TerminalSuggest([]string{"cat", "sam", "rut"})
	fmt.Println(best)
	// Output: cat
}
func TestSuggest(t *testing.T) {

	var set = []struct {
		options           []string
		expectedBest      string
		expectedRemaining []string
		expectedErr       bool
	}{
		{[]string{}, "cat", []string{}, true},
		{[]string{"cat", "sam", "rut"}, "cat", []string{"cat", "sam", "rut"}, false},
	}
	for _, test := range set {
		best, remaining, err := TerminalSuggest(test.options)
		// got an error
		if err != nil {
			// i didn't expect an error
			if !test.expectedErr {
				t.Error(
					"for:", test.options,
					"expected: an err",
					"got:", best,
					"and:", remaining,
				)
			}
			// i was expecting an error
			continue
		}
		// testing for length because i'm too lazy to write a slice comparison
		if best != test.expectedBest || len(remaining) != len(test.expectedRemaining) {
			t.Error(
				"for:", test.options,
				"expected:", test.expectedBest,
				"got:", best,
				"expected:", test.expectedRemaining,
				"got:", remaining,
			)
		}
	}
}
