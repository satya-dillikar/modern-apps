package main

import (
	"fmt"
	"strconv"
	"sync"
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
	producerCount      int
	consumerCount      int
	msgConsumedCount   int
	msgProducedCount   int
	producerWG         sync.WaitGroup
	consumerWG         sync.WaitGroup
	producerQueueIndex int
	consumerQueueIndex int
	jobsQueue          []string
	mutex              sync.Mutex
	producerGomutex    sync.Mutex
	consumerGomutex    sync.Mutex
	// producerGoWG       sync.WaitGroup
	// consumerGoWG       sync.WaitGroup
	jobs []string
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
	m.jobsQueue = make([]string, len(m.jobs))
	m.consumerQueueIndex = 0
	m.producerQueueIndex = 0
	// m.producerGoWG = sync.WaitGroup{}
	// m.consumerGoWG = sync.WaitGroup{}
	// m.producerGoWG.Add(len(m.jobs))
	// m.consumerGoWG.Add(0)
	m.consumerGomutex.Lock()
	//m.producerGomutex.Unlock()
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
	fmt.Printf("producerQueueIndex %v, consumerQueueIndex %v\n",
		m.producerQueueIndex, m.consumerQueueIndex)
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
	runToComplete := true
	for runToComplete {
		//Get a space for storing products
		m.producerGomutex.Lock()
		//Occupy product buffer
		m.mutex.Lock()

		i := m.producerQueueIndex
		msg := m.jobs[i]
		tempjob := "produced-by-" + strconv.Itoa(id) + ":" + msg
		m.producerQueueIndex = (m.producerQueueIndex + 1) % len(m.jobs)
		m.jobsQueue[m.producerQueueIndex] = tempjob
		m.msgProducedCount++
		if m.msgProducedCount > len(m.jobs) {
			runToComplete = false
		}
		//Release product buffer
		m.mutex.Unlock()
		//Inform consumers that there is a product
		m.consumerGomutex.Unlock()

	}
	//m.print("producer-" + strconv.Itoa(id))
}

func (m *multiProducerMultiConsumer) startProducers() {
	for i := 0; i < m.producerCount; i++ {
		m.producerWG.Add(1)
		go m.producer(i)
	}
}

func (m *multiProducerMultiConsumer) consumer(id int) {
	defer m.consumerWG.Done()
	runToComplete := true
	for runToComplete {
		//Must have products to consume
		m.consumerGomutex.Lock()
		//Lock buffer
		m.mutex.Lock()
		i := m.consumerQueueIndex
		b := m.jobsQueue[i]
		m.consumerQueueIndex = (m.consumerQueueIndex + 1) % len(m.jobs)
		m.msgConsumedCount++
		fmt.Printf("consumer-%v, mesg : %s\n", id, b)
		//Release buffer
		if m.msgConsumedCount > len(m.jobs) {
			runToComplete = false
		}
		m.mutex.Unlock()
		//Inform producers that there is room
		m.producerGomutex.Unlock()
	}
	//m.print("consumer-" + strconv.Itoa(id))
}

func (m *multiProducerMultiConsumer) startConsumers() {
	for i := 0; i < m.consumerCount; i++ {
		m.consumerWG.Add(1)
		go m.consumer(i)
	}
}

func main() {

	var m multiProducerMultiConsumer
	m.init(4, 2, messages)
	m.startProducers()
	m.startConsumers()
	m.cleanup()
}
