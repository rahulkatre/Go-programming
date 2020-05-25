package main

import "fmt"
//size of an array also included in type of that array, below array is defined to zero value i.e, 0 is stored in every
//byte
var buffer [256]byte
var buffer_1 [128]byte
// slice has type []byte called slice of bytes
var slice_2 []byte
// most idiomatic way writting slice
var slice_3 = buffer[0:12]

func main(){
	//here "slice_1" is an piece of an array
	slice_1 := buffer[128:130]
	//Below stmt is wrong since piece of an array sigifies slice. so slice data-structure should be used for below
	//operation
	//buffer_1 = buffer[1:12]
	slice_2 = buffer[0:12]

	//fmt.Println(buffer)
	fmt.Println(slice_1)
	fmt.Println(slice_2)
	fmt.Printf("%T\n", slice_1)
	fmt.Printf("%T\n", buffer)

}
