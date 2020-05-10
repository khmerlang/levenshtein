package levenshtein

func Distance(str1, str2 string) int {
	// var cost, lastdiag, olddiag int
	s1 := []rune(str1)
	s2 := []rune(str2)
	insertion_cost := 1
	deletion_cost := 1
	sub_cost := 1

	len_s1 := len(s1)
	len_s2 := len(s2)

	dist := make([][]int, len_s1+1)
	for i := range dist {
		dist[i] = make([]int, len_s2+1)
	}

	for x := 1; x <= len_s1; x++ {
		dist[x][0] = x
	}
	for y := 1; y <= len_s2; y++ {
		dist[0][y] = y
	}

	for x := 1; x <= len_s1; x++ {
		for y := 1; y <= len_s2; y++ {
			if s1[x-1] == s2[y-1] {
				dist[x][y] = dist[x-1][y-1]
			} else {
				dist[x][y] = minimum(dist[x-1][y]+deletion_cost, dist[x-1][y-1]+sub_cost, dist[x][y-1]+insertion_cost)
			}
		}
	}

	return (dist[len_s1][len_s2])
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
