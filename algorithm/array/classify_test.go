package array

import (
	"math/rand"
	"testing"
	"time"
)

/*
将一个数组进行分类，比如数组里面有奇数和偶数，将奇数放在数组前段，偶数放在后段，并且中间不改变奇偶元素间的相对顺序，如下:
input:  {3, 6, -5, 2, 8, -4, 7, 0, -1, 9, -9}
output: {3, -5, 7, -1, 9, -9, 6, 2, 8, -4, 0}
*/

/*
O(n^2) 复杂度的算法无需额外辅助空间，代码见 classifyON2
O(nlogn) 复杂度算法采用分治思想，但因为递归调用，会需要 logn 的额外辅助空间

当数组长度不是足够大时，classifyON2 的性能更好些；但当数组十分大，比如百万量级时，显然 classifyONLogn 性能更好
*/

func TestReverse(t *testing.T) {
	array := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	exp := []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}
	reverse(array)
	if exp, act := exp, array; !compareSlice(exp, act) {
		t.Errorf("got %v expected %v", act, exp)
	}
}

func TestClassify(t *testing.T) {
	array := []interface{}{3, 6, -5, 2, 8, -4, 7, 0, -1, 9, -9}
	plusMinusExp := []interface{}{3, 6, 2, 8, 7, 9, -5, -4, 0, -1, -9}
	oddEvenExp := []interface{}{3, -5, 7, -1, 9, -9, 6, 2, 8, -4, 0}

	testClassifyON2(t, plusMinusExp, copyArray(array), isPositive)
	testClassifyON2(t, oddEvenExp, copyArray(array), isOdd)

	testClassifyONLogn(t, plusMinusExp, copyArray(array), isPositive)
	testClassifyONLogn(t, oddEvenExp, copyArray(array), isOdd)
}

func testClassifyON2(t *testing.T, exp, array []interface{}, condition func(a interface{}) bool) {
	classifyON2(array, condition)
	if !compareSlice(exp, array) {
		t.Errorf("got %v expected %v", array, exp)
	}
}

func testClassifyONLogn(t *testing.T, exp, array []interface{}, condition func(a interface{}) bool) {
	classifyONLogn(array, condition)
	if !compareSlice(exp, array) {
		t.Errorf("got %v expected %v", array, exp)
	}
}

func compareSlice(exp, act []interface{}) bool {
	expLen, actLen := len(exp), len(act)
	if expLen != actLen {
		return false
	}
	minLen := expLen
	if actLen < minLen {
		minLen = actLen
	}

	for i := 0; i < minLen; i++ {
		if exp[i] != act[i] {
			return false
		}
	}
	return true
}

func copyArray(array []interface{}) []interface{} {
	copy := make([]interface{}, len(array), len(array))
	for i := 0; i < len(array); i++ {
		copy[i] = array[i]
	}
	return copy
}

func BenchmarkClassifyON2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 100000
	array := make([]interface{}, size, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Int()
	}
	for i := 0; i < b.N; i++ {
		classifyON2(array, isOdd)
	}
}

func BenchmarkClassifyONLogn(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 100000
	array := make([]interface{}, size, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Int()
	}
	for i := 0; i < b.N; i++ {
		classifyONLogn(array, isOdd)
	}
}
