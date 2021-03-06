package main 

import (
"fmt"

)
func PlusTwo() func(int) int{
    return func(x int) int {return x+2}
}

//use closures
func PlusX(x int) func(int) int{
    return func(y int) int{return x+y}
}

func PrintFunctionReturnsFunction(){
    p2 := PlusTwo()
    fmt.Printf("%v\n",p2(20))
    p20 := PlusX(20)
    fmt.Printf("%v\n",p20(20))    
}
