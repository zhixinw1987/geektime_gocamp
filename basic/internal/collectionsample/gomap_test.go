package collectionsample

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestStart(t *testing.T) {
	fmt.Printf("Map test start ... \n")
}

func TestCreateMap(t *testing.T) {
	m1, m2, m3, m4 := CreateMap()
	fmt.Printf("m1: %v, len(%d), size(%d), addr(%p)\n", m1, len(m1), unsafe.Sizeof(m1), &m1)
	fmt.Printf("m2: %v, len(%d), size(%d), addr(%p)\n", m2, len(m2), unsafe.Sizeof(m2), &m2)
	fmt.Printf("m3: %v, len(%d), size(%d), addr(%p)\n", m3, len(m3), unsafe.Sizeof(m3), &m3)
	fmt.Printf("m4: %v, len(%d), size(%d), addr(%p)\n", m4, len(m4), unsafe.Sizeof(m4), &m4)
	if m3 == nil {
		//m3["key"]="1"    panic here, m3 is nil, not initialized yet
		fmt.Println("m3 is nil")
	}
	m4["key"] = "val1" //m4 has been initialized with empty elements
	fmt.Printf("m4: %v, len(%d), addr(%p)\n", m4, len(m4), &m4)
}

func TestMapUpdate(t *testing.T) {
	_, _, _, m := CreateMap()

	fmt.Println("key is unique in map, assign value to same key will override value:")
	m["key1"] = "val1"
	fmt.Printf("m put key1: %v, len(%d), size(%d), addr(%p)\n", m, len(m), unsafe.Sizeof(m), &m)
	m["key1"] = "v1"
	fmt.Printf("m put key1: %v, len(%d), size(%d), addr(%p)\n\n", m, len(m), unsafe.Sizeof(m), &m)

	fmt.Println("find key from map, will always return default zero value if key not found")
	fmt.Println("with comma ok syntax, can check if the key actually found or not")
	val := FindKey(m, "kkk")
	fmt.Println("value found from kkk: " + val)
	val = FindOk(m, "kkk")
	fmt.Println("value found from kkk: " + val)

	fmt.Println("\ndelete from map, no failure or error if key not found in map")
	delete(m, "kkk")

	fmt.Println("map value stored in the bucket, when expansion occurs value will be migrated to new bucket, hence not possible to get pointer from map value &v")
	// fmt.Printf("%p\n", &m["abc"])

}

func TestIterateMap(t *testing.T) {
	fmt.Println("test map interation")
	_, m, _, _ := CreateMap()
	Coordinate(m).Iterate()
}
