package select

//simple concurrecny with goroutines with channel
//channels are ways for goroutines to communicate with each other
import (
	"fmt"
	"time"
)

func main5(){
	c1 := make(chan string)
	c2 := make(chan string)
	

	go func(){
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func(){
		for {
			c2 <- "every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()
	
	for {
		//select allows the channel to unblock
		select {
		case msg1 := <- c1:
				fmt.Println(msg1)
		}
		case msg2 := <- c2:
				fmt.Println(msg2)
		}
	}
}

