package array

import (
	"math/rand"
	"time"
)

func majorityON2(array []interface{}) interface{} {
	l := len(array)
	if l <= 0 {
		return nil
	}

	index := 0
	for index < l-1 {
		if array[index] != array[index+1] {
			index += 2
			continue
		}
		search := index + 2
		for ; search < l; search++ {
			if array[search] != array[index+1] {
				array[index+1], array[search] = array[search], array[index+1]
				index += 2
				break
			}
		}

		if search >= l {
			break
		}
	}

	return array[index]
}

func majorityON(array []interface{}) interface{} {
	l := len(array)
	if l <= 0 {
		return nil
	}

	index := 0
	search := 0
	for index < l-1 {
		if array[index] != array[index+1] {
			index += 2
			continue
		}

		if search <= index {
			search = index + 2
		}
		for ; search < l; search++ {
			if array[search] != array[index+1] {
				array[index+1], array[search] = array[search], array[index+1]
				index += 2
				break
			}
		}

		if search >= l {
			break
		}
	}

	return array[index]
}

func majorityO1(array []interface{}) interface{} {
	l := len(array)
	if l <= 0 {
		return nil
	}

	sample := make([]interface{}, 101, 101)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(sample); i++ {
		sample[i] = array[rand.Int()%l]
	}

	return majorityON(sample)
}
