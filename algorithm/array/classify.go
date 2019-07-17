package array

func leftSwamp(from, to int, array []interface{}) {
	tmp := array[to]
	for i := to - 1; i >= from; i-- {
		array[i+1] = array[i]
	}
	array[from] = tmp
}

func rightSwamp(from, to int, array []interface{}) {
	tmp := array[from]
	for i := from; i < to; i++ {
		array[i] = array[i+1]
	}
	array[to] = tmp
}

func classifyON2(array []interface{}, condition func(a interface{}) bool) {
	if len(array) < 2 {
		return
	}

	index := 0
	for search := index; search < len(array); search++ {
		if condition(array[search]) {
			leftSwamp(index, search, array)
			index++
		}
	}
}

func classifyONLogn(array []interface{}, condition func(a interface{}) bool) {
	l := len(array)
	if l < 2 {
		return
	}

	m := l / 2
	classifyONLogn(array[:m], condition)
	classifyONLogn(array[m:], condition)
	lIndex, rIndex := 0, m
	for ; lIndex < m && condition(array[lIndex]); lIndex++ {
	}
	for ; rIndex < l && condition(array[rIndex]); rIndex++ {
	}
	if lIndex >= m {
		return
	}

	reverse(array[lIndex:m])
	reverse(array[m:rIndex])
	reverse(array[lIndex:rIndex])
}

func isOdd(a interface{}) bool {
	return a.(int)&0x01 == 0x01
}

func isPositive(a interface{}) bool {
	return a.(int) > 0
}

func reverse(array []interface{}) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
