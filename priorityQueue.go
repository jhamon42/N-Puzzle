package main

// "fmt"

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	puz := x.(*puzzle)
	puz.index = n
	*pq = append(*pq, puz)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	puz := old[n-1]
	old[n-1] = nil
	puz.index = -1
	*pq = old[0 : n-1]
	return puz
}
