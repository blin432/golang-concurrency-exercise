package waitgroup

//simple concurrecny with goroutines with wait group
import (
	"fmt"
	"time"
	"sync"
)

func main2(){
	//init wait group
	var wg sync.WaitGroup
	//increment by one saying I have one go routine to wait for
	wg.Add(1)
	//invoking funciton
	go func(){
		count2("sheep")
		//decrements counter by one 
		wg.Done()
	}()
	//this will block until counter is at 0 ( no more goroutines)
	wg.Wait()
}

func count2(thing string){
	for i :=1; i<=5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}