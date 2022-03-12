package p1c1

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

var glbConsumedCount int = 0
var glbProducedCount int = 0

func producer(id int, link chan<- string) {
	for _, m := range messages {
		link <- "produced-by-" + strconv.Itoa(id) + ":" + m
		glbProducedCount++
	}
	close(link)
}

func consumer(id int, link <-chan string, done chan<- bool) {
	for b := range link {
		glbConsumedCount++
		fmt.Printf("consumer-%v, mesg : %s\n", id, b)
	}
	done <- true
}

func P1c1_main() {
	fmt.Println("No of messages:", len(messages))
	fmt.Printf("START: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)
	link := make(chan string)
	done := make(chan bool)
	go producer(1, link)
	go consumer(1, link, done)
	<-done
	fmt.Printf("END: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)
}
