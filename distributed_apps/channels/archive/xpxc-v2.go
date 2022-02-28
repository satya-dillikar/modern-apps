package main

import (
	"fmt"
	"strconv"
)

var messages = []string{
	"01-msg",
	"02-msg",
	"03-msg",
	"04-msg",
	"05-msg",
	"06-msg",
	"07-msg",
	"08-msg",
	"09-msg",
	"10-msg",
	"11-msg",
	"12-msg",
	"13-msg",
	"14-msg",
	"15-msg",
	"16-msg",
	"17-msg",
	"18-msg",
	"19-msg",
}

const producerCount int = 4
const consumerCount int = 2

var glbConsumedCount int = 0
var glbProducedCount int = 0

type producers struct {
	myQ  chan string
	quit chan bool
	id   int
}

func execute(jobQ chan<- string, workers []*producers, workerPool chan *producers, allDone chan<- bool) {
	for _, j := range messages {
		jobQ <- j
	}
	close(jobQ)
	for _, w := range workers {
		w.quit <- true
	}
	//close(workerPool)
	allDone <- true
}

func produce(jobQ <-chan string, p *producers, workerPool chan *producers) {
	for {
		select {
		case msg := <-jobQ:
			{
				if len(msg) > 0 {
					workerPool <- p
					glbProducedCount++
					p.myQ <- "produced-by-" + strconv.Itoa(p.id) + ":" + msg
					//fmt.Printf("Job \"%v\" produced by worker %v\n", msg, p.id)
				}
			}
		case <-p.quit:
			return
		}
	}
}

func consume(id int, workerPool <-chan *producers) {
	for {
		worker := <-workerPool
		if msg, ok := <-worker.myQ; ok {
			if len(msg) > 0 {
				glbConsumedCount++
				fmt.Printf("consumer-%v, mesg : %s\n", id, msg)
				//fmt.Printf("Message \"%v\" is consumed by consumer %v from worker %v\n", msg, cIdx, worker.id)
			}
		}
	}
}

func main() {
	var workers []*producers

	jobQ := make(chan string)

	allDone := make(chan bool)
	workerPool := make(chan *producers)

	fmt.Println("No of messages:", len(messages))
	fmt.Printf("START: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)

	for i := 0; i < producerCount; i++ {
		workers = append(workers, &producers{
			myQ:  make(chan string),
			quit: make(chan bool),
			id:   i,
		})
		go produce(jobQ, workers[i], workerPool)
	}

	go execute(jobQ, workers, workerPool, allDone)

	for i := 0; i < consumerCount; i++ {
		go consume(i, workerPool)
	}
	<-allDone
	close(workerPool)
	fmt.Printf("END: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)

}
