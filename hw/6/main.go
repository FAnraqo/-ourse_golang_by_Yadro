package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Couter struct {
	mutex sync.Mutex
	count int
}

func (n *Couter) Inc() {
	n.mutex.Lock()
	n.count++
	n.mutex.Unlock()
}

func (n *Couter) GetCount() int {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	return n.count
}

func main() {
	counter := &Couter{}

	var wg sync.WaitGroup
	for i := 0; i < rand.Intn(5)+5; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			fmt.Printf("Человет прошёл турникет. Прошло на данный момент: %d\n", counter.GetCount())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Итого людей прошло: %d\n", counter.GetCount())
}
