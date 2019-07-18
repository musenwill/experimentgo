package array

import (
	"math/rand"
	"testing"
	"time"
)

/*
若数组中某个元素出现次数比所有元素数量的一半还多，则这个元素是 majority 元素，那么如何找到这个 majority 元素呢？

1. 蛮力法就是统计第一个元素出现的次数，如果没有过半数，则接着统计第二个元素出现的次数。时间复杂度是 O(N ^ 2)
2. 可以对数组先排序，那么只需要对排序过的数组做一次遍历即可，时间复杂度是 O(N * logN)
3. 如果可以有 O(N) 的辅助空间，那么显然可以用 map 来分别记录每个元素出现的次数，时间复杂度是 O(N)，空间复杂度是 O(N)

4.
该题的条件是如此之强，上述方案显然达不到要求，理想的算法是不借助辅助空间，时间复杂度是 O(N)
根据该数组的性质，可以采取配对法，即每次移出两个不同的元素，直到找不到两个不同的元素了，那么最后剩下的元素必然是 majority。

该算法的第一个版本是 majorityON2。实际上，在数组随机性较强时，基本上处于 O(N) 的水平。但当算法随机性不够，导致某个元素会连续出现，
该连续串甚至会很长，这样 search 指针会额外带来 O(N) 的开销。那么在最坏情况下时间复杂度能达到 O(N ^ 2)

实际上 search 指针带来的 O(N) 开销完全可以避免，只要让 search 指针不回溯即可，那么 search 指针总共只会对数组遍历一次。这样
算法时间复杂度就是 O(N)。这就是第二个版本 majorityON

5. 最后有没有时间复杂度为 O(1) 的算法呢? 可以采取取样法，比如从数组中随机独立取样 100 个元素组成一个小数组，再用小数组去进行判断。
这样时间复杂度是 O(1)，尽管会存在小概率错误。但基本上只要样本足够大，错误会小到忽略不计。

*/

type mcase struct {
	array    []interface{}
	majority interface{}
}

func TestMajority(t *testing.T) {
	cases := []mcase{
		{
			array:    []interface{}{},
			majority: nil,
		},
		{
			array:    []interface{}{1},
			majority: 1,
		},
		{
			array:    []interface{}{1, 1},
			majority: 1,
		},
		{
			array:    []interface{}{1, 1, 1},
			majority: 1,
		},
		{
			array:    []interface{}{1, 2, 3, 4, 1, 1, 1, 1},
			majority: 1,
		},
		{
			array:    []interface{}{0, 1, 0, 1, 0, 1, 0},
			majority: 0,
		},
		{
			array:    []interface{}{2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			majority: 1,
		},
	}
	for _, c := range cases {
		runMajorityCaseON2(t, c)
		runMajorityCaseON(t, c)
		runMajorityCaseO1(t, c)
	}
}

func runMajorityCaseON2(t *testing.T, c mcase) {
	majority := majorityON2(c.array)
	if exp, act := c.majority, majority; exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
}

func runMajorityCaseON(t *testing.T, c mcase) {
	majority := majorityON(c.array)
	if exp, act := c.majority, majority; exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
}

func runMajorityCaseO1(t *testing.T, c mcase) {
	majority := majorityO1(c.array)
	if exp, act := c.majority, majority; exp != act {
		t.Errorf("got %v expected %v", act, exp)
	}
}

func BenchmarkMajorityON2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	candidate := []interface{}{0, 0, 0, 0, 1, 2, 3}
	size := 100000
	array := make([]interface{}, size, size)
	for i := 0; i < size; i++ {
		array[i] = candidate[rand.Int()%len(candidate)]
	}

	major := majorityON2(array)
	if exp, act := 0, major; exp != act {
		b.Errorf("got %v expected %v", act, exp)
	}
	for i := 0; i < b.N; i++ {
		majorityON2(array)
	}
}

func BenchmarkMajorityON(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	candidate := []interface{}{0, 0, 0, 0, 1, 2, 3}
	size := 100000
	array := make([]interface{}, size, size)
	for i := 0; i < size; i++ {
		array[i] = candidate[rand.Int()%len(candidate)]
	}

	major := majorityON(array)
	if exp, act := 0, major; exp != act {
		b.Errorf("got %v expected %v", act, exp)
	}
	for i := 0; i < b.N; i++ {
		majorityON(array)
	}
}

func BenchmarkMajorityO1(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	majority := 0
	size := 100000
	array := make([]interface{}, size, size)
	for i := 0; i < size; i++ {
		if rand.Float32() < 0.51 {
			array[i] = majority
		} else {
			array[i] = rand.Int()
		}
	}

	major := majorityO1(array)
	if exp, act := majority, major; exp != act {
		b.Errorf("got %v expected %v", act, exp)
	}
	for i := 0; i < b.N; i++ {
		majorityO1(array)
	}
}
