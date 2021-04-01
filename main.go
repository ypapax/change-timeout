package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("started")
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	ctx2, _ := context.WithTimeout(ctx, 5 * time.Second)
	<- ctx2.Done()
	log.Println("done")
}