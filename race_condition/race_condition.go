package main

import "sync"

func main() {

	/* Mutex prevents multiple goroutines from modifying shared data simultaneously.

		 A race condition occurs when multiple goroutines access the same variable concurrently,
		 and at least one modifies it without synchronization.

		 1.RWMutex

	Allows multiple readers but only one writer.

	Read :

	rw.RLock()

	fmt.Println(data)

	rw.RUnlock()

	write :

	rw.Lock()

	data++

	rw.Unlock()


	*/
	count := 0

	n := 1000

	var mu sync.Mutex

	var wg sync.WaitGroup

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()

		}()
	}
	wg.Wait()

	println("count", count)
}
