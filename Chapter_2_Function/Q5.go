package main 

import (
"fmt"

)

func PrintSliceAvg(floatSlice []float64) (avg float64){
    
    sum := 0.0
    switch len(floatSlice){
    case 0:
       avg = 0
    default: 
        for _,float := range floatSlice{
            sum +=  float
        }
        avg = sum/float64(len(floatSlice))
        fmt.Println(avg)
        }
    return
}
