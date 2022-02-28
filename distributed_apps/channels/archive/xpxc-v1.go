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

const producerCount int = 3
const consumerCount int = 2

var glbConsumedCount int = 0
var glbProducedCount int = 0

func msgLen(messages [][]string, pcnt int) int {
	lenCount := 0
	if pcnt > producerCount {
		return -1
	}
	for i := 0; i < producerCount; i++ {
		msgs := messages[i]
		lenCount += len(msgs)
	}
	return lenCount
}

func producer(id int, link chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := id * producerCount
	end := start + len(messages)/producerCount
	if end > len(messages) {
		end = len(messages)
	}
	for i := start; i < end; i++ {
		m := messages[i]
		link <- "produced-by-" + strconv.Itoa(id) + ":" + m
		glbProducedCount++
	}
}

func consumer(id int, link <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for b := range link {
		glbConsumedCount++
		fmt.Printf("consumer-%v, mesg : %s\n", id, b)
	}
}

func main() {
	fmt.Println("No of messages:", len(messages))
	fmt.Printf("START: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)

	link := make(chan string)
	wp := sync.WaitGroup{}
	wc := sync.WaitGroup{}

	for i := 0; i < producerCount; i++ {
		wp.Add(1)
		go producer(i, link, &wp)
	}
	for i := 0; i < producerCount; i++ {
		wc.Add(1)
		go consumer(i, link, &wc)
	}
	wp.Wait()
	close(link)
	wc.Wait()
	fmt.Printf("END: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)
}
