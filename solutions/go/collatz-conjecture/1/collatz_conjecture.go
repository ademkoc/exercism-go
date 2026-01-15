package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var result []int

	if n <= 0 {
		return 0, errors.New("n must be a positive integer")
	}

	for i := n; i != 1; {
		if i%2 == 0 {
			i = i / 2
		} else {
			i = i*3 + 1
		}
		result = append(result, i)
	}

	return len(result), nil
}
