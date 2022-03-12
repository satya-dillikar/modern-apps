package pxcx_v3

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
const consumerCount int = 5

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

func execute(jobQ chan<- string) {
	for _, j := range messages {
		jobQ <- j
	}
	close(jobQ)
}

func producer(jobQ <-chan string, quit chan bool, id int, link chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case msg := <-jobQ:
			{
				if len(msg) > 0 {
					link <- "produced-by-" + strconv.Itoa(id) + ":" + msg
					glbProducedCount++
				}
			}
		case <-quit:
			return
		}
	}
}

func consumer(id int, link <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for b := range link {
		glbConsumedCount++
		fmt.Printf("consumer-%v, mesg : %s\n", id, b)
	}
}

func Pxcx_v3_main() {
	//fmt.Println("No of messages:", msgLen(messages, producerCount))
	fmt.Println("No of messages:", len(messages))

	fmt.Printf("START: glbConsumedCount:%v, glbProducedCount:%v\n",
		glbConsumedCount, glbProducedCount)

	jobQ := make(chan string)
	link := make(chan string)
	wp := sync.WaitGroup{}
	wc := sync.WaitGroup{}

	for i := 0; i < producerCount; i++ {
		wp.Add(1)
		go producer(jobQ, i, link, &wp)
	}
	go execute(jobQ)
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
