package mergesort

func Mergesort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	// split functionality
	split := func(list []int) ([]int, []int) {
		mid := len(list) >> 1

		return list[:mid], list[mid:]
	}

	// merge functionality
	res := []int{}
	a, b := split(list)
	i, j := 0, 0

	for {
		if i >= len(a) {
			res = append(res, b...)
			break
		}
		if j >= len(b) {
			res = append(res, a...)
			break
		}

		if a[i] > b[j] {
			res = append(res, b[j])
			j++
		} else {
			res = append(res, a[i])
			i++
		}
	}

	return res
}
