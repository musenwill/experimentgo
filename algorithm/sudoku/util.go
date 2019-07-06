package sudoku

func unique(list []int) []int {
	set := make(map[int]struct{})
	u := make([]int, 0, len(list))

	for _, val := range list {
		if _, ok := set[val]; !ok {
			set[val] = struct{}{}
			u = append(u, val)
		}
	}

	return u
}