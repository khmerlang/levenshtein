package levenshtein

func Distance(str1, str2 string) int {
	var cost, lastdiag, olddiag int
	s1 := []rune(str1)
	s2 := []rune(str2)

	len_s1 := len(s1)
	len_s2 := len(s2)
	insertion_cost := 1
	deletion_cost := 1
	sub_cost := 2

	dist := make([]int, len_s1+1)

	for y := 1; y <= len_s1; y++ {
		dist[y] = y
	}

	for x := 1; x <= len_s2; x++ {
		dist[0] = x
		lastdiag = x - 1
		for y := 1; y <= len_s1; y++ {
			olddiag = dist[y]
			cost = 0
			if s1[y-1] != s2[x-1] {
				cost = sub_cost // substitution cost
			}

			dist[y] = minimum(dist[y]+deletion_cost, dist[y-1]+insertion_cost, lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return dist[len_s1]
}

func minimum(num_1, num_2, num_3 int) int {
	min := num_1
	if min > num_2 {
		min = num_2
	}

	if min > num_3 {
		min = num_3
	}

	return min
}
