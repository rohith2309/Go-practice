package main

import "fmt"

func main(){

fmt.Println(add(5,6))
fmt.Println(add(5,4))

fmt.Println(sums(10,20,30,40))

nums:=[]int{1,3,4,5,6}

fmt.Println(sums(nums...))
oddgen:=nextodd(1)
fmt.Println(oddgen())
fmt.Println(oddgen())
fmt.Println(oddgen())

}

func sums(nums ...int)int{

	ans:=0

	for _,n:=range(nums){
		ans+=n
	}
	fmt.Println("the sum is :")
	return ans
}

func add(x int,y int)(int,string){
    ans:=x+y
	var str string=""

	if ans>10{
		str="greater than 10 "
	}else
	{
	   str="less than 10"
	}

	return x+y,str

}


//closures


func nextodd(x int) func()int{

	

	return func() int {
		x+=2
		return x
	}
 }