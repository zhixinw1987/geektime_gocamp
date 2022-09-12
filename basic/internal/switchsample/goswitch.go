package switchsample

import "fmt"

type Animal interface {
	say()
}

type Cat struct {
	Name string
}
type Dog struct {
	Name string
}

func (cat *Cat) say() {
	fmt.Println("Cat says meow")
}
func (dog *Dog) say() {
	fmt.Println("Dog says warf")
}

func SwtichExecOrder(x int) {
	fmt.Println(`switch execution sequence, from top to buttom, left to right
like if clause, place highest hit case on top left whenever possible to optimize performance
with keyword fallthrough, next case will also be executed and skip the evaluation of the condition for next case
no keyword break in each case, once any first matched case executed, the switch process ends, will not evaluate any further case even if any matched`)
	switch initSwitchExpr(x) {
	case checkCase1():
		fmt.Println("case 1 matched")
	case checkCase2_1(), checkCase2_2():
		fmt.Println("case 2 matched")
	case checkCase3():
		fmt.Println("case 3 matched")
		fmt.Println("")
		fallthrough
	case checkCase4():
		fmt.Println("case 4 matched")
	case checkCase_3():
		fmt.Println("case _3 matched")
	default:
		fmt.Println("default")
	}
}

func SwitchType(x any, a Animal) {
	fmt.Println(".(type) is special syntax in switch clause to exam the type of given variable")
	var tx interface{} = x
	switch tx.(type) {
	case int, uint:
		fmt.Println("found int")
	case rune, string:
		fmt.Println("found word")
	case float32, complex128:
		fmt.Println("found float")
	default:
		fmt.Println("unidentified")
	}

	fmt.Println("if a custom interface is evaluated in switch, all cases must be the concreate type of the interface")
	switch t := a.(type) {
	case *Cat:
		fmt.Printf("it's cat %v\n", t)
	case *Dog:
		fmt.Printf("it's dog %v\n", t)
	default:
		fmt.Println("unidentified")
	}
}

func initSwitchExpr(x int) int {
	fmt.Println("init switch expression")
	return x
}

func checkCase1() int {
	fmt.Println("eval case 1")
	return 1
}

func checkCase2_1() int {
	fmt.Println("eval case 2_1")
	return 0
}

func checkCase2_2() int {
	fmt.Println("eval case 2_2")
	return 2
}

func checkCase3() int {
	fmt.Println("eval case 3")
	return 3
}

func checkCase_3() int {
	fmt.Println("eval case _3")
	return 3
}

func checkCase4() int {
	fmt.Println("eval case4")
	return 4
}
