package main
import "fmt"

func main(){
    elements := make(map[string]string)
    elements["H"] ="Hydrogen"
    elements["He"] ="Helium"
    elements["Li"] ="Lithium"
    elements["Be"] ="Beryllium"
    elements["B"] ="Boron"
    elements["C"] ="Carbon"
    elements["N"] ="Nitrogen"
    elements["O"] ="Oxygen"
    elements["F"] ="Fluorine"
    elements["Ne"] ="Heon"

//    fmt.Println(elements["Li"])
//    fmt.Println(elements["Un"])

    if name,ok := elements["Li"]; ok {
        fmt.Println(name, ok)
    }

    if name1,ok1 := elements["Un"]; ok1{
        fmt.Println(name1, ok1)
    }else{
        fmt.Println("Failed to find")
    }


}
