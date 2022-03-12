package main

import (
	"fmt"

	"satya.com/producer_consumer/p1c1"
	"satya.com/producer_consumer/p1cx"
	"satya.com/producer_consumer/pxc1"
	"satya.com/producer_consumer/pxcx"
)

func main() {

	fmt.Println("\n*** SINGLE PRODUCER SINGLE CONSUMER ***")
	fmt.Println("Calling P1c1_main ...")
	p1c1.P1c1_main()

	fmt.Println("\n*** SINGLE PRODUCER MULTIPLE CONSUMER ***")
	fmt.Println("Calling P1cx_main ...")
	p1cx.P1cx_main()

	fmt.Println("\n*** MULTIPLE PRODUCER SINGLE CONSUMER ***")
	fmt.Println("Calling Pxc1_main ...")
	pxc1.Pxc1_main()

	fmt.Println("\n*** MULTIPLE PRODUCER MULTIPLE CONSUMER ***")
	fmt.Println("Calling Pxcx_main ...")
	pxcx.Pxcx_main()

	fmt.Println("\nDONE.")
}
