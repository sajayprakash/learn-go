package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	// var <variable_name> <data_type> = <value>
	var intNum int = 69
	fmt.Println(intNum)

	// float64 - 64-bit floating point numbers
	// float32 - 32-bit floating point numbers
	// float32 is less precise than float64, but takes up less memory space and is faster to process
	var float32Num float32 = 12345678.9
	var float64Num float64 = 12345678.9
	fmt.Println(float32Num)
	fmt.Println(float64Num)

	// var result = intNum + floatNum - Cannot add int and float

 	// Convert intNum to float64
	var result = float64(intNum) + float64Num
	fmt.Println(result)

	var myString string = "Hello World!"
	var myString2 string = "Γ"
	fmt.Println(myString)
	fmt.Println(myString2)

	// len() gives the no of bytes in a string
	fmt.Println(len(myString)) 

	// len() does not work for non-ASCII characters
	fmt.Println(len(myString2)) 

	// utf8.RuneCountInString() gives the no of characters in a non-ASCII string
	fmt.Println(utf8.RuneCountInString(myString2))
	 
	// Rune is different from string
	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	// Infer data type
	var inferedInt = 69
	fmt.Println(inferedInt)

	// Short-hand declaration
	shortHandInt := 69
	fmt.Println(shortHandInt)

	// Short-hand declaration with multiple variables
	shortHandInt1, shortHandInt2 := 69, 420
	fmt.Println(shortHandInt1, shortHandInt2)

	//Don't do this, we cannot tell the data type of the variable from this line
	whatever1 := foo() 
	fmt.Println(whatever1)

	// Do this instead to make the code more readable
	var whatever2 string = foo() 
	fmt.Println(whatever2)

	const myConst = 69
	fmt.Println(myConst)

	// myConst = 420 - Cannot assign to myConst
}

// func <function_name>(<parameter_name> <data_type>) <return_type> {
	// return <value>
// }
func foo() string {
	return "Hello World!"
}