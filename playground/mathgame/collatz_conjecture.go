package mathgame

func collatzConjecture(n int) int {
	if n%2 == 0 {
		return n / 2
	}

	return 3*n + 1
}

func g(n int) int {
	if n == 2 {
		return 1
	}
	return 1 + g(collatzConjecture(n))
}
