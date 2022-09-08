package collectionsample

import (
	"fmt"
)

type Postion struct {
	x float64
	y float64
}

type Coordinate map[Postion]string

//m1, m2 map created with initial value
//m3 declared map as nil
//m4 declared map with no value, but initial size of 10
func CreateMap() (map[int][]string, map[Postion]string, map[string]string, map[string]string) {
	//m := make(map[[]int]string)
	//m := make(map[func(){}]string)
	//m := make(map[map[int]int]string)
	//map key must support == and != operand,
	//in above samples, slice, func, map can only compare to nil, can't be used as map key

	m1 := map[int][]string{
		1: {"val1", "val2"},
		2: {"v1", "v2"},
	}

	m2 := map[Postion]string{
		{29.9348, 52.3174}: "hospital",
		{33.1153, 76.9831}: "school",
	}

	var m3 map[string]string
	m4 := make(map[string]string, 10)
	return m1, m2, m3, m4
}

func FindKey(m map[string]string, k string) string {
	return m[k]
}

func FindOk(m map[string]string, k string) string {
	v, ok := m[k]
	if ok {
		return v
	} else {
		return "k not found"
	}
}

func (c Coordinate) Iterate() {
	for k, v := range c {
		fmt.Printf("coordinate [%v] is location [%s]\n", k, v)
	}
}
