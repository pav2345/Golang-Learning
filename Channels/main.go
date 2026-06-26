package main 

func main(){
	var ch  chan int // declaration of channel , It is nil not instanttiated

	ch = make(chan int)

	ch<- 100 // sending  the value 
	v:=<-ch
	ch<-200

	println(v)

	//-->=<=

	


}


//communication go uses channels 
//channels are queues
//flow of data from one direction to other it can never bidirectional 
//why we need channels?
//what kind of channels?
//sender and receiver 
//the Size of channel is very important when sender sends the data  only if channel  buffer is free
//Block --->
//the sender is Blocked until the Receiver recives the  value
//the Reciver is blocked until sender send the value
// the size channel is one element