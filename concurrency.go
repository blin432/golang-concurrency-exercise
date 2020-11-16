package concurrency

//simple concurrecny with goroutines
import (
	"fmt"
	"time"
)

func main(){
	// adding "go" creates a goroutines to run concurrently. If you run without go it does not go to next line "fish"
	//if you add go to create a goroutine  for fish, nothing will show because it is looking for next line after "fish" so the next program is done and terminates
	//can solve this problem with "time.Sleep(time.Second *) or use fmt.Scanln() because it is a blocking call "
	go count("sheep")
	count("fish")
}

func count(thing string){
	for i :=1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}