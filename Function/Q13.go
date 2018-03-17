package main 

import (
"fmt"

)

func BubbleSort(l []int){
    for i:=0;i<len(l);i++ {
        for j:=i+1;j<len(l);j++{
            if l[j]<l[i] {
                l[j], l[i] = l[i],l[j]
            }
        }
    }
}

func PrintBubbleSort(){
    
    a := []int{2,100,19,6,32,82}
    fmt.Printf("unsorted: %v\n",a)
    BubbleSort(a)
    fmt.Printf("sorted: %v\n",a)
}
