package main

import (
	"fmt"
	"time"
)

func main() {
	var toSort int
	c := make(chan int)
	for {
		a, err := fmt.Scan(&toSort)
		if a == 0 || err != nil {
			break
		}
		go func(){
            sleepSort(toSort)
            c <- toSort
        }()
	}
    <- c
	fmt.Printf("\n")
}

func sleepSort(toSort int) {
	napTime := time.Millisecond * time.Duration(toSort)
	time.Sleep(napTime)
    fmt.Printf("%d ", toSort)
}
