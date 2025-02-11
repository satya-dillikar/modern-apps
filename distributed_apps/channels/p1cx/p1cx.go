package p1cx

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

const consumerCount int = 3

var glbConsumedCount int = 0
var glbProducedCount int = 0

func producer(id int, link chan<- string) {
	for _, m := range messages {
		link <- "produced-by-" + strconv.Itoa(id) + ":" + m
		glbProducedCount++
	}
	close(link)
}

func consumer(id int, link <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for b := range link {
		glbConsumedCount++
		fmt.Printf("consumer-%v, mesg : %s\n", id, b)
	}
}

func P1cx_main() {
	fmt.Println("No of messages:", len(messages))
	fmt.Printf("START: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)
	link := make(chan string)
	//done := make(chan bool)
	wg := sync.WaitGroup{}

	go producer(1, link)
	for i := 0; i < consumerCount; i++ {
		wg.Add(1)
		go consumer(i, link, &wg)
	}
	wg.Wait()
	//<-done
	fmt.Printf("END: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)
}
