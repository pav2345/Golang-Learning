package main 


import (
	"fmt"
	"time"
	"sync"

)

func main(){

	wg:=new(sync.waitGroup) //state =0
	wg.Add(1)
	go func(){
		for i:=1 ; i<=50; i++{
			time.Sleeptime.Millisecond + Time 

		}

	}
}



// similar to join handle ---> wait Group 
//the Sync wait group are primitive standard Library in go concurrency that lets you wait  for multiple goroutines to finish before continuing Execution 
