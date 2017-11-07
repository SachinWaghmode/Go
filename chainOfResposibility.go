package main

import "fmt"

type Handler interface{
	handleRequest(request string)
	SetSuccessor(next Handler)
}

type ConcreteHandler1 struct{
	Successor Handler
}

func (handler1 ConcreteHandler1) handleRequest(request string) {
 	fmt.Println("R1 got the request...") 
	if ( request == "R1" ){
            fmt.Println(" => This one is mine!")
        }else{
	     fmt.Println(" => This one is not mine!")
            	if ( handler1.Successor != nil ){
			fmt.Println("Test!")
			handler1.Successor.handleRequest(request);
		}	      
        }
}

func (handler1 *ConcreteHandler1) SetSuccessor(next Handler) {
        handler1.Successor = next 
	}

type ConcreteHandler2 struct{
	Successor Handler
}

func (handler2 ConcreteHandler2) handleRequest(request string) {
 	fmt.Println("R2 got the request...") 
	if ( request == "R2" ){
            fmt.Println(" => This one is mine!")
        }else{
	     fmt.Println(" => This one is not mine!")
            	if ( handler2.Successor != nil ){
			fmt.Println("Test!")
			handler2.Successor.handleRequest(request);
		}	
                
        }
}

func (handler2 *ConcreteHandler2) SetSuccessor(next Handler) {
        handler2.Successor = next 
	}


type ConcreteHandler3 struct{
	Successor Handler
}

func (handler3 ConcreteHandler3) handleRequest(request string) {
 	fmt.Println("R3 got the request...") 
	if ( request == "R3" ){
            fmt.Println(" => This one is mine!")
        }else{
	     fmt.Println(" => This one is not mine!")
            	if ( handler3.Successor != nil ){
			fmt.Println("Test!")
			handler3.Successor.handleRequest(request);
		}	
                
        }
}

func (handler3 *ConcreteHandler3) SetSuccessor(next Handler) {
        handler3.Successor = next 
	}


func Test(){
	h1 := new(ConcreteHandler1)
	h2 := new(ConcreteHandler2)
	h3 := new(ConcreteHandler3)
	
	h1.SetSuccessor(h2)
	h2.SetSuccessor(h3)
	
	fmt.Println("Sending R2...")
	h1.handleRequest("R2")
	
	fmt.Println("Sending R3...")
	h1.handleRequest("R3")
	
	fmt.Println("Sending R1...")
	h1.handleRequest("R1")
	
	fmt.Println("Sending RX...")
	h1.handleRequest("RX")
}

func main(){
	Test()
}
