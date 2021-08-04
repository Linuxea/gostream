package slice

import (
	"fmt"
	"strconv"
	"testing"
)

var s = []T{1, 2, 3, 4, 5, 6, 7, 9, 8, 7, 6, 5, 4, 3, 2, 1}

func TestForeach(t *testing.T) {

	foreach(s, func(t T) {
		fmt.Println(t)
	})
}

func TestReduce(t *testing.T) {

	result := reduce(s, 0, func(t1, t2 T) T {
		return t1.(int)*10 + t2.(int)
	})
	fmt.Println(result)
}

func TestMapper(t *testing.T) {
	stringSlice := mapper(s, func(t T) T {
		return strconv.Itoa(t.(int))
	})
	fmt.Println(stringSlice)
}

func TestFilter(t *testing.T) {
	filterResult := filter(s, func(t T) bool {
		return t.(int)%2 == 0
	})
	fmt.Println(filterResult)
}

func TestAnyMatch(t *testing.T) {
	match := anyMatch(s, func(t T) bool {
		return t.(int) > 0
	})
	fmt.Println(match)
}

func TestAllMatch(t *testing.T) {
	match := allMatch(s, func(t T) bool {
		return t.(int)%2 == 0
	})
	fmt.Println(match)
}

func TestFirst(t *testing.T) {
	first := first(s, func(t T) bool {
		return t.(int) > 3
	})
	fmt.Println(first)
}

func TestLast(t *testing.T) {
	result := last(s, func(t T) bool {
		return t.(int) >= 9
	})
	fmt.Println(result)
}

func TestSort(t *testing.T) {
	ts := sort(s, func(t1, t2 T) bool {
		return t1.(int) >= t2.(int)
	})
	fmt.Println(ts)
}

func TestCompare(t *testing.T) {
	ts := compare(s, func(t1, t2 T) bool {
		return t1.(int) > t2.(int)
	})
	fmt.Println(ts)
}
