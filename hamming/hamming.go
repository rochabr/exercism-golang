package hamming

import (
	"fmt"
)

//Distance should have a comment
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("len A != len B")
	}

	i, sum := 0, 0
	for i < len(a) {
		if a[i] != b[i] {
			sum++
		}
		i++
	}

	return sum, nil
}
