package constantsample

import (
	"fmt"
	"strconv"
)

type myInt int

const a = 10
const b myInt = 20

//enum iota start from 0, increment by 1
const (
	SUNDAY = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

const (
	_         = iota //0: assign iota to empty value, skip 0 in enum
	NORTH            //1
	SOUTH            //2
	EAST             //3
	WEST             //4
	_                //5: assign to empty, skip number in enum
	NORTHEAST = "NE" //NE: assign custom value
	NORTHWEST        //NE: auto repeat previous value
	SOUTHEAST = iota //8: assign back iota incremental value
	SOUTHWEST        //9
)

type Gender int

const (
	MALE Gender = iota //define the enum type as Gender
	FEMALE
	TRANSGENDER
	UNKNOW
)

//toString function
func (gen Gender) String() string {
	switch gen {
	case 0:
		return "男"
	case 1:
		return "女"
	case 2:
		return "变性"
	default:
		return "未知"
	}
}

func constAdd() {
	fmt.Println(a + b)
	fmt.Println(`go const can be decalred without type, actual type will be automatically assumed
	not like vars, const can cast type implicitly as long the types are compatible
	a is no type const, compile will assume the type by initial value
	a + b is implicit cast, type myInt is extend of type int, hence can be cast successfully
	`)
}

func simpleEnum() {
	fmt.Println("weekday is " + strconv.Itoa(FRIDAY))
	fmt.Println(EAST, WEST, NORTHEAST, NORTHWEST, SOUTHEAST, SOUTHWEST)
}

func extendedEnum() {
	fmt.Println(MALE)
}
