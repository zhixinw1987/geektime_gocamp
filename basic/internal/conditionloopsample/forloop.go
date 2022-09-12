package conditionloopsample

import (
	"fmt"
	"strconv"
)

func forloops() {
	fmt.Println("simple loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("infinite loop")

	i := 0
	for {
		fmt.Println("infinite loop" + strconv.Itoa((i)))
		i++
		if i > 10 {
			break
		}
	}

	fmt.Println("for range")
	sl := []int{1, 2, 3, 4, 5}
	for _, v := range sl {
		fmt.Println("range of " + strconv.Itoa(v))
	}
}

func rangeIsCopy() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	fmt.Println("before loop: ", a)
	fmt.Println("position of a[0]", &a[0])
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
			fmt.Printf("position of v [%p], of a[%p]\n", &v, &a[0])
		}
		r[i] = v
	}
	fmt.Println(a)
	fmt.Println(r)
	fmt.Println(`in the range loop, is manipulating the copy of original array
		for i,v := range a ==> for i,v := range a'
	`)

	a = [5]int{1, 2, 3, 4, 5}
	r = [5]int{}
	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
			fmt.Printf("position of v [%p], of a[%p]\n", &v, &a[0])
		}
		r[i] = v
	}
	fmt.Println("after loop by pointer")
	fmt.Println(a)
	fmt.Println(r)
}
