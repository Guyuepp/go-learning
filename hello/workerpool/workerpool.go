package workerpool

import (
	"fmt"
	"math/rand/v2"
)

type Job struct {
	id  int
	num int
}

type Result struct {
	job *Job
	res int
}

func Test() {
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	createPool(64, jobChan, resultChan)
	go func() {
		for {
			res := <-resultChan
			fmt.Println(*res.job, res.res)
		}
	}()
	id := 0
	for {
		job := &Job{id, rand.Int()}
		jobChan <- job
		id++
	}
}

func createPool(num int, jobChan <-chan *Job, resultChan chan<- *Result) {
	for range num {
		go func(jobChan <-chan *Job, resultChan chan<- *Result) {
			for job := range jobChan {
				res := 0
				num := job.num
				for num > 0 {
					res += num % 10
					num /= 10
				}
				resultChan <- &Result{job, res}
			}
		}(jobChan, resultChan)
	}
}
