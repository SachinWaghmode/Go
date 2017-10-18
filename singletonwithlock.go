package main

import (
	"fmt";
	"sync"
	
)

type singleton struct{
}

var instance *singleton
var mu sync.Mutex

func getInstance() *singleton {
    if instance == nil {     
        mu.Lock()
        defer mu.Unlock()

        if instance == nil {
            instance = &singleton{}
        }
    }
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