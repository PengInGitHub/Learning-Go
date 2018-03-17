package main 

import (
"fmt"

)

func Map(f func (int) int , l []int) []int{
    j := make([]int, len(l))
    for k,v := range l {
        j[k] = f(v)
    }
    return j
}

func PrintMap(){
    f := func(i int) int{
        return i*i
    }
    m := []int{7,8,9}
    mapResults := Map(f,m)
    
    fmt.Println(mapResults)
}
