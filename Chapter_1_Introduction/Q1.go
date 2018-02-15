package main 

import (
"fmt"

)

//Step 1
func PrintForLoop(){
    for i:=0;i<10;i++ {
        fmt.Println(i)
    }
}

//Step 2
func PrintGotoLoop(){
    
    i:=0
    Loop:
        fmt.Println(i)
    if i<10 {
        i++
        goto Loop
    }
}