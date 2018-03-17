package main 

import (
"fmt"

)
func Min(l []int) (min int){
    min = l[0]
    for _,v := range l {
        if v<min {
            min = v
        }
    } 
    return
}

func Max(l []int) (max int){
    max = l[0]
    for _,v := range l{
        if v>max {
            max = v
        }
    }
    return
}

func PrintMinMax(){
    a := []int{2,7,1,89,45}
    fmt.Println(Max(a))    
    fmt.Println(Min(a))
}
