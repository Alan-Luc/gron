package main

import (
	"fmt"
	"time"

	"github.com/Alan-Luc/gron/scheduler"
)

func test() {
	fmt.Println("this is a test")
}

func main() {
	s := &scheduler.Scheduler{}

	for len(s.Jobs) > 0 {
		s.RunNextJob()
		time.Sleep(time.Second)
	}
}
