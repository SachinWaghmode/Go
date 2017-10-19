package main

import (
	"fmt";
	"sync"
)
type singleton struct{
}

var instance *singleton

var once sync.Once

func getInstance() *singleton{
	once.Do(func() {
	        if (instance == nil){
			instance =&singleton{}
		}
	})
	return instance 
}

func Test(){
	i1 := getInstance();
	i2 := getInstance();

	if i1 == i2 {
		fmt.Println("instance of i1 and i2 are  equal")
	}
} 

func main(){
	Test()
}