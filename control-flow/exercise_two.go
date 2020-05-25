package main

import "fmt"

func main() {

	for i:=65; i<=90; i++{
		for j:=0; j<3; j++{
			fmt.Printf("%#U\n", i)
		}
	}
	// If expression not provided to swith keyword, it wil by default evaluates for true value
	switch {
	case 1==1:
		fmt.Println("Equal")
		// fallthrough lets others cases to be executed no matter whether the cases passes or fails.
		fallthrough
	case  !true:
		fmt.Println("Not Equal")
	default:
		fmt.Println("Default case")
	}
}
