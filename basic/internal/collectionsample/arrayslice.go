package collectionsample

import (
	"fmt"
	"strconv"
	"unsafe"
)

func InitArray() ([5]uint8, [6]uint8, [3]string, [4]uint8) {
	var arr1 [5]uint8 // 0 0 0 0 0
	var arr2 = [6]uint8{1, 2, 3, 4, 5, 6}
	var arr3 = [...]string{"a", "b", "c"} //auto detect array len
	var arr4 = [4]uint8{3: 10}            //0 0 0 10  -> assign init value to given pos
	return arr1, arr2, arr3, arr4
}

func InitSlice() ([]uint8, []uint8, []uint8) {
	sli1 := make([]uint8, 5, 10) //uint 8 slice, init len 5, cap 10
	var empSl1 []uint8
	var empSl2 = []uint8{}
	return sli1, empSl1, empSl2
}

func CreateSliceFromArray(arr *[6]uint8) ([]uint8, []uint8) {
	//arr [1,2,3,4,5,6]
	sl1 := arr[:4:5]                //1,2,3,4 cap(5)=max(5)-low(0)
	sl2 := arr[3:len(arr):cap(arr)] //4,5,6 cap(3)=max(6)-low(3)
	return sl1, sl2
}

func PeekArray(arr [5]uint8) {
	fmt.Println(arr[0])
}

func CheckSize(arr [3]string) {
	fmt.Println("element count: " + strconv.Itoa(len(arr)))
	fmt.Println("mem bytes: " + fmt.Sprint(unsafe.Sizeof(arr)))
}
