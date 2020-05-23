package main

import "fmt"

/*for x assign 42
for y assign “James Bond”
for z assign true
*/
var x = 42
var y = "James Bond"
var z = true
//var s string

func main(){

	//fmt.Sprintf("%d %s %t\n", x,y,z)
	s := fmt.Sprintf("%d %s %t\n", x,y,z)
	fmt.Printf("%s", s)

}
