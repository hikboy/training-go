
package main
import "fmt"

func main(){
    var x [5]float64
    x[0]=66
    x[1]=6
    x[2]=26
    x[3]=46

    var total float64
    for _, value:=range x {
        total += value
    }

    fmt.Println(total)

}
