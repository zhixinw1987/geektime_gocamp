package collectionsample

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr1, arr2, arr3, arr4 := InitArray()
	fmt.Printf("arr1: %v, len(%d), cap(%d), addr(%p)\n", arr1, len(arr1), cap(arr1), &arr1)
	fmt.Printf("arr2: %v, len(%d), cap(%d), addr(%p)\n", arr2, len(arr2), cap(arr2), &arr2)
	fmt.Printf("arr3: %v, len(%d), cap(%d), addr(%p)\n", arr3, len(arr3), cap(arr3), &arr3)
	fmt.Printf("arr4: %v, len(%d), cap(%d), addr(%p)\n", arr4, len(arr4), cap(arr4), &arr4)
	CheckSize(arr3)
	fmt.Printf("test array completed \n\n\n")
}

func TestSlice(t *testing.T) {
	_, arr, _, _ := InitArray()

	_, initSl2, initSl3 := InitSlice()
	fmt.Printf("emptyArr1: %v, len(%d), cap(%d), addr(%p)\n", initSl2, len(initSl2), cap(initSl2), &initSl2)
	fmt.Printf("emptyArr2: %v, len(%d), cap(%d), addr(%p)\n", initSl3, len(initSl3), cap(initSl3), &initSl3)
	fmt.Println(`both arrays are empty, functionally equivalent
	emptyArr1 is nil, while emptyArr2 is initialized as zero length array
	the nil slice is the preferred style`)

	sl1, sl2 := CreateSliceFromArray(&arr)

	// sl1 := arr[:4:5]                //1,2,3,4 cap(5)=max(5)-low(0)
	// sl2 := arr[3:len(arr):cap(arr)] //4,5,6 cap(3)=max(6)-low(3)

	fmt.Printf("sli1: %v, len(%d), cap(%d), addr(%p)\n", sl1, len(sl1), cap(sl1), &sl1)
	fmt.Printf("sli2: %v, len(%d), cap(%d), addr(%p)\n", sl2, len(sl2), cap(sl2), &sl2)
	fmt.Printf("================slice created from array=====================\n\n")

	fmt.Println(`Slice is a handler to underlying array 
	any modification to the element in slice will also be reflected in original array
	also if another slice is created from the same array, any change will be reflected on same element in both slice`)

	sl1[3] = 100
	fmt.Printf("arr: %v, len(%d), cap(%d), addr(%p)\n", arr, len(arr), cap(arr), &arr)
	fmt.Printf("sli1: %v, len(%d), cap(%d), addr(%p)\n", sl1, len(sl1), cap(sl1), &sl1)
	fmt.Printf("sli2: %v, len(%d), cap(%d), addr(%p)\n", sl2, len(sl2), cap(sl2), &sl2)
	fmt.Printf("====================slice modification================================\n\n")

	sl1 = append(sl1, 10)
	sl1 = append(sl1, 20)
	sl1[3] = 110
	fmt.Printf(`once slice expanded cap, the underlying array has been copied to a new array with extended cap
	any change to the slice element, will no long reflect in original array, but apply to the new array with extended cap`)
	fmt.Printf("arr: %v, len(%d), cap(%d), addr(%p)\n", arr, len(arr), cap(arr), &arr)
	fmt.Printf("sli1: %v, len(%d), cap(%d), addr(%p)\n", sl1, len(sl1), cap(sl1), &sl1)
	fmt.Printf("sli2: %v, len(%d), cap(%d), addr(%p)\n", sl2, len(sl2), cap(sl2), &sl2)

}
