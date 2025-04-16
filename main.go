package main

import (
	"fmt"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	stream := generator(done, 10000)
	pipeline := add(done, square(done, stream))

	fmt.Println(<-pipeline)
}
