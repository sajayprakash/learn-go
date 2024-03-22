package main

import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println("Hello World!")
	var result,remainder,err  =intDivision(420, 69)
	if err!=nil {
		fmt.Println("Error:",err)
	}else if remainder == 0 {
		fmt.Printf("The result is: %d\n",result)
	}else {
		fmt.Printf("The result is: %d and the remainder is: %d\n",result,remainder)
	}

	// Conditional Switch statement
	switch remainder {
	case 0:
		fmt.Println("The remainder is 0")
	case 1:
		fmt.Println("The remainder is 1")
	default:
		fmt.Println("The remainder is neither 0 nor 1")
	}
}

// A function can return multiple values
func intDivision(a int, b int) (int,int,error) {
	var err error
	if b == 0 {
		err = errors.New("cannot divide by zero")
		return 0,0,err
	}
	var result int = a/b
	var remainder int = a%b
	return result,remainder,err
}