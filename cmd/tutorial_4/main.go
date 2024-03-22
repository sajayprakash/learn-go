package main

import "fmt"

func main(){
	// var myArray [5]int32 = [5]int32{1,2,3,4,5}

	// ... is used to infer the size of the array from the number of elements
	myArray := [...]int32{1,2,3,4,5}
	fmt.Println(myArray[0])
	fmt.Println(myArray[2:5])

	// & is used to get the memory address of a variable
	// Arrays are stored in contiguous memory locations
	fmt.Println(&myArray[2])
	fmt.Println(&myArray[3])

	// Slices are a reference to an array and are not stored in contiguous memory locations
	var intSlice []int = []int{1,2,3,4,5}

	// cap() gives the capacity of the slice and len() gives the length of the slice
	fmt.Println(intSlice, cap(intSlice), len(intSlice))

	intSlice = append(intSlice, 6)
	fmt.Println(intSlice, cap(intSlice), len(intSlice))

	var intSlice2 []int = []int{7,8,9}

	// ... is also used to append one slice to another
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice, cap(intSlice), len(intSlice))

	// make() is used to create a slice with a specified capacity
	// Slices that have pre-allocated memory are faster than slices that don't
	var intSlice3 []int32 = make([]int32, 3,8)
	fmt.Println(intSlice3, cap(intSlice3), len(intSlice3))

	// Maps are key-value pairs
	var map1 map[string]uint8 = make(map[string]uint8)
	map1["a"] = 1
	fmt.Println(map1)

	var map2 map[string]uint8 = map[string]uint8{"a":1, "b":2, "c":3, "d":4}
	fmt.Println(map2)

	// Maps always return a value and a boolean
	value,ok:=map2["e"]
	if ok {
		fmt.Println(value)
	}else {
		fmt.Println("Key not found")
	}

	// Delete a key-value pair
	delete(map2, "b")
	fmt.Println(map2)

	// Iterate over a map
	// Ordered is not preserved
	for key,value := range map2 {
		fmt.Println(key,value)
	}

	// For loops
	var i,j int = 0,0

	for i<5 {
		fmt.Println(i)
		i++
	}

	for {
		if j>10 {
			break
		}
		fmt.Println(j)
		j++
	}

	for k:=2;k<8;k++ {
		fmt.Println(k)
	}
}