package main

import (
	"fmt"
	"sync"
	"strconv"
	"os"
	// "math/rand"
)

type Node struct {
	id int
	to chan int
	from chan int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func node_process(node Node, wg *sync.WaitGroup) {
	to := node.to
	from := node.from

	receive := <- from
	if receive == -1 {
		to <- node.id
		master := <- from
		fmt.Println(master)
		defer wg.Done()
	}	else {
		max := max(node.id, receive)
		to <- max
	}
}

func main() {
	random_num, _ := strconv.Atoi(os.Args[1])
	var wg sync.WaitGroup
	total_nodes := 1000000
	channels := make(map[int]chan int)

	for i := 0; i < total_nodes; i++ {
		channels[i] = make(chan int)
	}

	var node Node
	for i := 1; i < total_nodes; i++ {
		node = Node{i, channels[i], channels[i-1]}
		go node_process(node, &wg)
	}
	node = Node{0, channels[0], channels[total_nodes-1]}
	wg.Add(1)
	go node_process(node, &wg)

	channels[random_num] <- -1
	wg.Wait()
}
