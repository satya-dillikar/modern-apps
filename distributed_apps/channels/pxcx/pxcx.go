package pxcx

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
	producerCount    int
	consumerCount    int
	msgConsumedCount int
	msgProducedCount int
	producerWG       sync.WaitGroup
	consumerWG       sync.WaitGroup
	jobsQueue        chan string
	jobs             []string
}

func (m *multiProducerMultiConsumer) init(producerCount int, consumerCount int,
	jobs []string) {
	m.producerCount = producerCount
	m.consumerCount = consumerCount
	m.msgConsumedCount = 0
	m.msgProducedCount = 0
	m.jobsQueue = make(chan string)
	m.producerWG = sync.WaitGroup{}
	m.consumerWG = sync.WaitGroup{}
	m.jobs = jobs
	fmt.Println("No of jobs:", len(m.jobs))
	fmt.Printf("START: msgConsumedCount:%v, msgProducedCount:%v\n",
		m.msgConsumedCount, m.msgProducedCount)
}

func (m *multiProducerMultiConsumer) cleanup() {

	m.producerWG.Wait()
	close(m.jobsQueue)
	m.consumerWG.Wait()
	fmt.Printf("END: msgConsumedCount:%v, msgProducedCount:%v\n",
		m.msgConsumedCount, m.msgProducedCount)
}

func (m *multiProducerMultiConsumer) producer(id int) {
	defer m.producerWG.Done()
	start := id * m.producerCount
	end := start + len(m.jobs)/m.producerCount
	if end > len(m.jobs) {
		end = len(m.jobs)
	}
	for i := start; i < end; i++ {
		msg := m.jobs[i]
		m.jobsQueue <- "produced-by-" + strconv.Itoa(id) + ":" + msg
		m.msgProducedCount++
	}
}

func (m *multiProducerMultiConsumer) startProducers() {
	for i := 0; i < m.producerCount; i++ {
		m.producerWG.Add(1)
		go m.producer(i)
	}
}

func (m *multiProducerMultiConsumer) consumer(id int) {
	defer m.consumerWG.Done()
	for b := range m.jobsQueue {
		m.msgConsumedCount++
		fmt.Printf("consumer-%v, mesg : %s\n", id, b)
	}
}

func (m *multiProducerMultiConsumer) startConsumers() {
	for i := 0; i < m.consumerCount; i++ {
		m.consumerWG.Add(1)
		go m.consumer(i)
	}
}

func Pxcx_main() {

	var m multiProducerMultiConsumer
	m.init(3, 3, messages)
	m.startProducers()
	m.startConsumers()
	m.cleanup()
}
