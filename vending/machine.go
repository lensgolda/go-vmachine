package vending

import "fmt"

var (
	in    chan interface{} = make(chan interface{})
	out   chan string      = make(chan string)
	quite chan int         = make(chan int)
)

func HotDogMachineV1() (chan interface{}, chan string) {

	go func() {
		fmt.Printf("VMachine [IN]: got %v\n", <-in)
		out <- "hot dog"
	}()

	return in, out
}

func HotDogMachineV2(capacity int) (chan interface{}, chan string, chan int) {

	go func() {
		for capacity > 0 {
			input := <-in

			switch input.(type) {
			case string:
				out <- fmt.Sprintf("Expected type int, got %T\n", input)
			case int:
				if input == 3 {
					out <- "hot dog\n"
					capacity--
				} else {
					out <- "wilted letuce\n"
				}
			default:
				out <- fmt.Sprintf("Unexpected input of type %T\n", input)
			}
		}
		quite <- 0
	}()

	return in, out, quite
}
