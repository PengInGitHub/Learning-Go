package main 
import "fmt"

func decalreOutsieloop(){
    
    var i int
    for i=0;i<10;i++{
        fmt.Println(i) 
    }
    fmt.Println(i) 
}
