package main 

import (
"fmt"

)

type stack struct{
    i int
    data [10]int
}

func(s *stack) push(k int){
    s.data[s.i] = k
    s.i++
}

func(s *stack) pop() int{
    s.i--
    return s.data[s.i]
}

func PrintStack(){
    var s stack
    s.push(25)
    fmt.Printf("stack:v%\n",s)
}
