package main 

import (
"fmt"

)

func PrintNumbers(numbers ... int){
    for _,d := range numbers{
        fmt.Println(d)
    }
}
