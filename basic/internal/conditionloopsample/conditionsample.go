package conditionloopsample

import "fmt"

var m map[string]string

func happyPathofIf() {
	fmt.Println(`in GO is recommended to apply happy path to the left edge, 
	exit early from the function if condition not matched,
	to avoid nested if checking,
	put happy return statement as the very last line, avoid else block if possible
	`)

	m = map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
	}

	if _, ok := m["key1"]; !ok {
		fmt.Println("key1 not found")
		return
	}
	v1 := m["key1"]
	if _, ok := m["key2"]; !ok {
		fmt.Println("key2 not found")
		return
	}
	v2 := m["key2"]
	if v, ok := m["key3"]; !ok {
		fmt.Println("key3 not found")
		return
	} else {
		fmt.Println("value of key3 is " + v)
	}
	v3 := m["key3"]
	fmt.Println(v1 + v2 + v3)
}
