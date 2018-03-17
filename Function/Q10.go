package main 

import (
"fmt"

)

func fibonacci(value int) []int{
    fibonacciArray := make([]int,value)
    fibonacciArray[0],fibonacciArray[1] = 1, 1
    for n:=2; n<value; n++ {
        fibonacciArray[n] = fibonacciArray[n-1] + fibonacciArray[n-2]
    }
    return fibonacciArray

}

func PrintFibonacci(){
    
    fmt.Println(fibonacci(50))
}
