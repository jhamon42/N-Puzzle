package main

// "fmt"

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
