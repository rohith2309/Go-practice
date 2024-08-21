package main

import (
	"fmt"
)

func main(){

	fmt.Println("hello range")

	num:=[]int{10,23,43,24,22,11}

	for n:=range num{
		fmt.Println("index %d",n)
	}

	for _,n:=range(num){
		fmt.Println("value %d",n)
	}
	
	mpp:=map[string]int{"a":1,"b":2,"v":22,"r":18}

	for k:=range(mpp){
		
		fmt.Println("key %d",k)
		
	}
	for k,v:=range(mpp){
		
		fmt.Println("key ",k,"value ",v)
		
	}
}	