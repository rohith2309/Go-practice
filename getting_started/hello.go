package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
    var name string
    fmt.Scan(&name)
    fmt.Println("your name is ",name)
	//fmt.Println("%d",Add(10,20))
    
    var arr[10] int
    //n:=2
    // fmt.Println("enter 2 numbers")
    // for i:=0;i<n;i++{
    //     fmt.Scan(&arr[i])

    // }

     slice:=make([]int,6)
     fmt.Println(slice)
     slice=append(slice,90,20,12)
     fmt.Println(len(slice))
    
     cp1:=make([]int,len(slice))
     copy(cp1,slice)
     cp1=append(cp1,42069)

     fmt.Println(slice)
     fmt.Println(cp1)
     fmt.Println(cp1[5:])

     //MAPS ==== map[key]value

     mpp:=make(map[string]int)

     mpp["k1"]=07
     mpp["k2"]=188

     val,present:=mpp["k"]
     fmt.Println(mpp)

     fmt.Println(present,val)
     delete(mpp,"k1")
     delete(mpp,"l")
     fmt.Println(mpp)


     


    for i:=0;i<2;i++{
        fmt.Println("i = ",i,"hello",name ,arr[i])
    }
}
