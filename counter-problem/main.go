package main 

import (
	"time"
	"fmt"
	"sync"
	"time"


)


var {
	Counter =0
	wg =new(sync.waitGroup)

}


func main(){
	for  range 100{
		wg.add(1)
		go Inc()
		go Dec()
	}

	println(Counter )
	runtime.Goexit()
}