package pxc1

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

const producerCount int = 4

var glbConsumedCount int = 0
var glbProducedCount int = 0

func producer(id int, link chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := id * producerCount
	end := start + len(messages)/producerCount
	if end > len(messages) {
		end = len(messages)
	}
	// fmt.Printf("producer %v, start %v, end %v\n",
	// 	id, start, end)
	//for _, m := range messages[id] {
	for i := start; i < end; i++ {
		m := messages[i]
		link <- "produced-by-" + strconv.Itoa(id) + ":" + m
		glbProducedCount++
	}
}

func consumer(id int, link <-chan string, done chan<- bool) {
	for b := range link {
		glbConsumedCount++
		fmt.Printf("consumer-%v, mesg : %s\n", id, b)
	}
	done <- true
}

func Pxc1_main() {
	fmt.Println("No of messages:", len(messages))
	fmt.Printf("START: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)
	link := make(chan string)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	for i := 0; i < producerCount; i++ {
		wg.Add(1)
		go producer(i, link, &wg)
	}
	go consumer(1, link, done)
	wg.Wait()
	close(link)
	<-done
	fmt.Printf("END: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)
}
