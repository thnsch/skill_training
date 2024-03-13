package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	philoNumber     int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(eatRequest, eatResponse chan bool, eatPermission chan int) {

	const shouldEat int = 3
	hasEaten := 0

	for hasEaten < shouldEat {
		eatRequest <- true // ask for allowance
		<-eatResponse      // wait for response

		p.leftCS.Lock()
		p.rightCS.Lock()

		hasEaten++
		fmt.Printf("Philosopher %d: Starting to eat (%d of %d).\n", p.philoNumber, hasEaten, shouldEat)

		time.Sleep(1000 * time.Millisecond)

		fmt.Printf("Philosopher %d: Finishing eating.\n", p.philoNumber)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-eatPermission // unregister the eating philosopher, saying it's finished
	}

	return
}

func host(p []*Philo, done chan bool) {

	eatRequest1 := make(chan bool)
	eatRequest2 := make(chan bool)
	eatRequest3 := make(chan bool)
	eatRequest4 := make(chan bool)
	eatRequest5 := make(chan bool)
	eatResponse1 := make(chan bool)
	eatResponse2 := make(chan bool)
	eatResponse3 := make(chan bool)
	eatResponse4 := make(chan bool)
	eatResponse5 := make(chan bool)
	eatPermission := make(chan int, 2)

	// philosophers start eating
	go p[0].eat(eatRequest1, eatResponse1, eatPermission)
	go p[1].eat(eatRequest2, eatResponse2, eatPermission)
	go p[2].eat(eatRequest3, eatResponse3, eatPermission)
	go p[3].eat(eatRequest4, eatResponse4, eatPermission)
	go p[4].eat(eatRequest5, eatResponse5, eatPermission)

	for {
		select {
		case <-eatRequest1:
			eatPermission <- 1   // register the eating philosopher
			eatResponse1 <- true // send allowance
		case <-eatRequest2:
			eatPermission <- 1
			eatResponse2 <- true
		case <-eatRequest3:
			eatPermission <- 1
			eatResponse3 <- true
		case <-eatRequest4:
			eatPermission <- 1
			eatResponse4 <- true
		case <-eatRequest5:
			eatPermission <- 1
			eatResponse5 <- true
		case <-time.After(3 * time.Second): // timeout when no incoming data
			fmt.Println("All finished.")
			close(eatRequest1)
			close(eatRequest2)
			close(eatRequest3)
			close(eatRequest4)
			close(eatRequest5)
			done <- true
			return
		}
	}

}

func main() {

	done := make(chan bool)

	// create chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// create philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	// start host goroutine
	go host(philos, done)

	<-done      // wait for data
	close(done) // close channel

}
