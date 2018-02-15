package main 

import (
"fmt"

)



func PrintFizzBuzz(){
    
    const(
        FIZZ = 3
        BUZZ = 5        
    )  
   
    var p bool //to judge wether to print the values
    
    for i:=0;i<100;i++{

         p = false

         if i%FIZZ==0 {
             fmt.Println("FIZZ")
             p = true
         }

         if i%BUZZ==0 {
             fmt.Println("BUZZ")
             p = true
         }

         if !p {
             fmt.Printf("%v",i)
         }

         fmt.Println()
    }
    
}
