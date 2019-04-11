package main

import "fmt"

func main(){
    var x[5]float64
    x[0]=64
    x[1]=65
    x[2]=97
    x[3]=23
    var total float64
    for i:=0;i<len(x);i++ {
        total += x[i]
    }
    fmt.Println(total)
    fmt.Println(len(x))
    fmt.Println(total / float64(len(x)))
}
