package workerpools

import "fmt"

//worker pools where there is a que of work to be done and multiple concurrent workers pulling workers off the que.

func main6() {
	//buffer channels
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//worker as a concurrent goroutine
	go worker(jobs, results)

	//we can add more workers by just adding workers
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	//populating channel and then showing results of channel
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}

}

// we are only receive from the jobs channel and send only on the results channel
func worker(jobs <-chan int, results chan<- int) {
	//que of numbers, use range feature to consume items in this que.
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
