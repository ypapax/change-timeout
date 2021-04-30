package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	timeouts(time.Second, 5 * time.Second)
	timeouts(5 * time.Second, time.Second)
	timeouts(0, 2 * time.Second)

}

func timeouts(t1, t2 time.Duration) {
	log.Println("started", t1, t2)
	var ctx context.Context
	if t1 == 0 {
		ctx = context.Background()
	} else {
		var cancel1 context.CancelFunc
		ctx, cancel1 = context.WithTimeout(context.Background(), t1)
		defer cancel1()
	}

	ctx2, cancel2 := context.WithTimeout(ctx, t2)
	defer cancel2()
	<- ctx2.Done()
	log.Println("done")
}