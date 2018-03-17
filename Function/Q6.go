package main 



func IntOrdering(int1, int2 int)(int, int){
    if int1 > int2 {
        return int2, int1
    }
    return int1, int2
}
