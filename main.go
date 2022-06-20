package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var fact = new(big.Int) //I declare a pointer to the type *big.Int
	fact.MulRange(1, 100)   //The MulRange function under the hood finds the factorial
	str := fact.String()    //I set two variables:
	var supSlice []int      //str - I convert fact to string, supSlice - I create a variable with an int slice
	for i, _ := range str { //since the string under the hood is a slice, it can be branched with for/range
		x := string(str[i])             //the x variable converts the i element of str slide from uint8 to string
		x1, _ := strconv.Atoi(x)        //variable x1 converts variable x from string to int
		supSlice = append(supSlice, x1) //use the append function to add the variable x1 to our created int slice
	} //go through the whole slice str (our number is 100!)
	answer := 0 //answer is a counter variable
	for _, v := range supSlice {
		answer += v //enumeration of the intrinsic slide in which
	} //the value of each element of that slide is added to the variable answer
	fmt.Println(answer)
}
