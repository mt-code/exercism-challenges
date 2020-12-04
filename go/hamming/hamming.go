package hamming

import "fmt"

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, fmt.Errorf("both parameters must be the same length")
	}

	hammingDistance := 0
	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
