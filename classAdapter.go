package main

import "fmt"

type Target interface {
	sayHello()
}

type AdapteeParent struct {
}


func (adaptee1 AdapteeParent) printMessage(msg string)  {
	fmt.Println(msg)
}


type Adapter struct {
	AdapteeParent 
}


func (receiver Adapter) sayHello() {	
	adapter := Adapter{AdapteeParent{}}
	adapter.printMessage("Hello")	
}


func Test(){
	ad := Adapter{AdapteeParent{}}
	ad.sayHello()
}

func main(){
	Test()
}