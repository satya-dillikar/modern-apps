package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"satya.com/concurrency/queue"
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

type multiProducerMultiConsumer struct {
	producerCount    int
	consumerCount    int
	msgConsumedCount int
	msgProducedCount int
	producerWG       sync.WaitGroup
	consumerWG       sync.WaitGroup
	//producerQueueIndex int
	jobsQueue *queue.CircleQueue
	jobs      []string
	mutex     sync.Mutex
}

func (m *multiProducerMultiConsumer) init(producerCount int, consumerCount int,
	jobs []string) {
	m.producerCount = producerCount
	m.consumerCount = consumerCount
	m.msgConsumedCount = 0
	m.msgProducedCount = 0
	//m.jobsQueue = make(chan string)
	m.producerWG = sync.WaitGroup{}
	m.consumerWG = sync.WaitGroup{}
	m.jobs = jobs
	//m.producerQueueIndex = 0
	m.jobsQueue, _ = queue.Init(len(m.jobs))
	fmt.Println("No of jobs:", len(m.jobs))
	fmt.Printf("START: msgConsumedCount:%v, msgProducedCount:%v\n",
		m.msgConsumedCount, m.msgProducedCount)
}

func (m *multiProducerMultiConsumer) print(name string) {
	fmt.Printf("----------START:%s----------\n", name)
	// fmt.Println("jobs:", m.jobs)
	fmt.Printf("producerCount %v, consumerCount %v\n",
		m.producerCount, m.consumerCount)
	fmt.Printf("msgProducedCount %v, msgConsumedCount %v\n",
		m.msgProducedCount, m.msgConsumedCount)
	// fmt.Println("jobsQueue:", m.jobsQueue)

	fmt.Printf("----------END:%s----------\n", name)

}
func (m *multiProducerMultiConsumer) cleanup() {

	m.producerWG.Wait()
	//close(m.jobsQueue)
	m.consumerWG.Wait()
	fmt.Printf("END: msgConsumedCount:%v, msgProducedCount:%v\n",
		m.msgConsumedCount, m.msgProducedCount)
}

func (m *multiProducerMultiConsumer) producer(id int) {
	defer m.producerWG.Done()
	for true {
		m.mutex.Lock()
		if m.msgProducedCount >= len(m.jobs) {
			m.mutex.Unlock()
			break
		}
		i := m.msgProducedCount
		msg := m.jobs[i]
		tempjob := "produced-by-" + strconv.Itoa(id) + ":" + msg
		m.jobsQueue.Push(tempjob)
		m.msgProducedCount++
		time.Sleep(1 * time.Second)
		m.mutex.Unlock()
	}
	fmt.Printf("producer id %v done\n", id)
}

func (m *multiProducerMultiConsumer) startProducers() {
	for i := 0; i < m.producerCount; i++ {
		m.producerWG.Add(1)
		go m.producer(i)
	}
}

func (m *multiProducerMultiConsumer) consumer(id int) {
	defer m.consumerWG.Done()
	for true {
		m.mutex.Lock()
		if m.msgConsumedCount >= len(m.jobs) {
			m.mutex.Unlock()
			break
		}
		b, err := m.jobsQueue.Pop()
		if err != nil {
			//fmt.Println(err)
			//continue
		} else {
			m.msgConsumedCount++
			fmt.Printf("consumer-%v, mesg : %s\n", id, b)
		}
		//time.Sleep(1 * time.Second)
		// fmt.Printf("msgConsumedCount %v MaxSize %v\n",
		// 	m.msgConsumedCount, m.jobsQueue.MaxSize)
		m.mutex.Unlock()
	}
	fmt.Printf("consumer id %v done\n", id)
}

func (m *multiProducerMultiConsumer) startConsumers() {
	for i := 0; i < m.consumerCount; i++ {
		m.consumerWG.Add(1)
		go m.consumer(i)
	}
}

func main() {

	var m multiProducerMultiConsumer
	m.init(3, 5, messages)
	m.startProducers()
	m.startConsumers()
	m.cleanup()
}
