package main 

import (
"fmt"

)
func PlusTwo() func(int) int{
    return func(x int) int {return x+2}
}

func PrintFunctionReturnsFunction(){
    p2 := PlusTwo()
    fmt.Printf("%v\n",p2(20))
}
