package main

import (
	"fmt"

	"github.com/lensgolda/vmachine/vending"
)

func main() {

	/* HotDog Machine Version 1 (simple check) */

	// in, out := vending.HotDogMachineV1()
	// in <- "pocket lint"
	// result := <-out
	// fmt.Printf("VMachine [OUT]: %v\n", result)

	/* HotDog Machine Version 2 */

	in, out, quite := vending.HotDogMachineV2(5)

	for {
		select {
		case in <- "pocket lint":
			v := <-out
			fmt.Printf("VMachine [OUT]: %v", v)
		case in <- 3:
			v := <-out
			fmt.Printf("VMachine [OUT]: %v", v)
		case in <- 2:
			fmt.Printf("VMachine [OUT]: %v", <-out)
		case <-quite:
			fmt.Println("Quit...")
			return
		}
	}
}
