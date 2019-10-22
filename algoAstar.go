package main

import (
	"container/heap"
	"fmt"
)

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push :
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	puz := x.(*Puzzle)
	puz.index = n
	*pq = append(*pq, puz)
}

// Pop :
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	puz := old[n-1]
	old[n-1] = nil
	puz.index = -1
	*pq = old[0 : n-1]
	return puz
}

func (nm PuzzleMap) get(p Puzzle) *Puzzle {
	n, ok := nm[p.label]
	if !ok {
		nm[p.label] = p
		return &p
	}
	return &n
}

func aStarAlgo(env *Env) error {
	puzMap := PuzzleMap{}
	puzQueue := &PriorityQueue{}
	heap.Init(puzQueue)
	startPuz := puzMap.get(env.initial)
	startPuz.open = true
	heap.Push(puzQueue, startPuz)
	for {
		if puzQueue.Len() == 0 {
			return nil
		}

		current := heap.Pop(puzQueue).(*Puzzle)

		current.open = false
		current.closed = true

		if current == puzMap.get(env.goal) {
			// Found a path to the goal.
			// p := []myMap{}
			// curr := current
			// for curr != nil {
			// 	p = append(p, curr.array)
			// 	curr = curr.parent
			// }
			return nil
		}
		// fmt.Println(current)
		fmt.Println()
		for _, neighbor := range current.findNeighbor(env) {
			fmt.Println(neighbor)

			if neighbor.puzMap == nil {
				continue
			}
			cost := current.g + 1
			neighborNode := puzMap.get(neighbor)
			if cost < neighborNode.g {
				if neighborNode.open {
					heap.Remove(puzQueue, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = true
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.open = true
				heap.Push(puzQueue, neighborNode)
			}
		}
	}
}
