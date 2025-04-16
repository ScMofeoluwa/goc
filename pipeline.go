package main

import "math/rand"

func generator(done <-chan interface{}, bufferSize int) <-chan int {
	stream := make(chan int, bufferSize)
	go func() {
		defer close(stream)
		for i := 0; i < bufferSize; i++ {
			select {
			case <-done:
				return
			case stream <- rand.Intn(100) + 1:
			}
		}
	}()
	return stream
}

func square(done <-chan interface{}, numStream <-chan int) <-chan int {
	stream := make(chan int)
	go func() {
		defer close(stream)
		for i := range numStream {
			select {
			case <-done:
				return
			case stream <- i * i:
			}
		}
	}()
	return stream
}

func add(done <-chan interface{}, numStream <-chan int) <-chan int {
	stream := make(chan int, 1)
	sum := 0
	go func() {
		defer close(stream)
		for i := range numStream {
			sum += i
		}
		select {
		case <-done:
			return
		case stream <- sum:
		}
	}()
	return stream
}
