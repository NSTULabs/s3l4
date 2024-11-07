package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosof struct {
	name  string
	left  chan struct{}
	right chan struct{}
}

func (p *Philosof) Eat() {
loop:
	for {
		select {
		case p.right <- struct{}{}:
			select {
			case p.left <- struct{}{}:
				break loop
			default:
				<-p.right
				time.Sleep(100 * time.Millisecond)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println(p.name, "eating")
	time.Sleep(time.Second) // eating
	fmt.Println(p.name, "stop eating")

	<-p.left
	<-p.right
}

func main() {
	names := []string{"sokrat", "platon", "aristotel", "gerodot", "foma_akvinskiy"}

	philosophers := make([]*Philosof, len(names))
	prev := make(chan struct{}, 1)
	first := prev
	for i, name := range names {
		if i != len(names)-1 {
			next := make(chan struct{}, 1)
			philosophers[i] = &Philosof{
				name:  name,
				left:  prev,
				right: next,
			}
			prev = next
		} else {
			philosophers[i] = &Philosof{
				name:  name,
				left:  prev,
				right: first,
			}
		}
	}

	var wg sync.WaitGroup
	for {
		for i := 0; i < len(philosophers); i++ {
			wg.Add(1)
			go func() {
				philosophers[i].Eat()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
