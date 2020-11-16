package channel

//simple concurrecny with goroutines with channel
//channels are ways for goroutines to communicate with each other
import (
	"fmt"
	"time"
	"sync"
)

func main3(){
	//pass channel in 
	c := make(chan string)
	go count3 "sheep", c)

	//receive message from channel
	//sending and receiving are blocking operations
	//when you are trying to receive something you have to wait until something is sent
	for msg := range c { //iterating over the rage of the channel
		fmt.Println(msg)
	}
}

func count3(thing string, c chan string){
	for i :=1; i<=5; i++ {
		c <- thing  //sending message
		time.Sleep(time.Millisecond * 500)
	}

	//closing channel
	close(c)
}