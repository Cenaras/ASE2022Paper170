package main

import "sync"

func main() {
	c := sync.Cond{
		L: &sync.Mutex{},
	}
	c.Wait()
}
