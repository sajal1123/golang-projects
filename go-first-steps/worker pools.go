package main

import (
	"fmt"
	"time"
)

func worker_func(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job ", j)
		results <- j * 2
	}
}
func main() {
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 0; w < 3; w++ {
		go worker_func(w, jobs, results)
	}

	for j := 0; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for r := 0; r <= numJobs; r++ {
		<-results
	}
}
