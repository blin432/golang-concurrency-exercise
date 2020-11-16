package channel2

//simple concurrecny with goroutines with channel
//channels are ways for goroutines to communicate with each other
import (
	"fmt"
)

func main4(){
	//making a buffered channel with a capacity
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	// adding third channel will not work because channel will be full
	//c <- "three"

	msg := <- c
	fmt.Println(msg)

	msg = <- c 
	fmt.Println(msg)

}

