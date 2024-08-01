package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	id     int
	number int
}

type result struct {
	job job
	res int
}

func sumOfDigits(number int) int {
	res := 0
	for number != 0 {
		digit := number % 10
		res += digit
		number /= 10
	}
	time.Sleep(2 * time.Second)
	return res
}

func worker(jobs chan job, results chan result, wg *sync.WaitGroup) {
	wg.Add(1)
	for job := range jobs {
		res := sumOfDigits(job.number)
		results <- result{job, res}
	}
	wg.Done()
}

func workerObserver(results chan result, wg *sync.WaitGroup) {
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int, jobs chan job) {
	for i := 0; i < noOfJobs; i++ {
		jobs <- job{i, rand.Intn(999)}
	}
	close(jobs)
}

func main() {
	var jobs = make(chan job, 3)
	var results = make(chan result, 5)
	var wg sync.WaitGroup

	starTime := time.Now()

	go allocate(100, jobs)

	for i := 0; i < 30; i++ {
		go worker(jobs, results, &wg)
	}

	go workerObserver(results, &wg)

	for result := range results {
		fmt.Printf("Job id %d, random number: %d. Sum of digits: %d\n", result.job.id, result.job.number, result.res)
	}

	endTime := time.Now()
	diff := endTime.Sub(starTime)
	fmt.Println("Total time taken ", diff.Seconds())
}
