package main 

import (
"fmt"

)

func PrintSliceAvg(floatSlice []float64){
    sum := 0.0
    for _,float := range floatSlice{
        sum +=  float
    }
    avg := sum/float64(len(floatSlice))
    fmt.Println(avg)
}
