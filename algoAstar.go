package main

import (
	"container/heap"
	"fmt"
)

func (nm PuzzleMap) get(p *Puzzle) *Puzzle {
	n, ok := nm[p.label]
	if !ok {
		nm[p.label] = p
		return p
	}
	return n
}

func aStarAlgo(env *Env) error {
	puzMap := PuzzleMap{}
	puzQueue := &PriorityQueue{}
	heap.Init(puzQueue)
	startPuz := puzMap.get(&env.initial)
	startPuz.open = true
	heap.Push(puzQueue, startPuz)
	// k := 0
	for {
		// * fin si la PriorityQueue est vide
		if puzQueue.Len() == 0 {
			return nil
		}

		// * on recup le top du top dans les open bien sur
		current := heap.Pop(puzQueue).(*Puzzle)
		current.open = false
		current.close = true
		// * fin si goal map
		if current == puzMap.get(&env.goal) {
			// Found a path to the goal.
			// p := []myMap{}
			// curr := current
			// for curr != nil {
			// 	p = append(p, curr.array)
			// 	curr = curr.parent
			// }
			fmt.Println("oui c'est un vrai GGGGGGG")
			fmt.Println(current.puzMap)
			fmt.Println()
			return nil
		}

		// ? c'est quoi l'addes du current
		// fmt.Printf("\n\ndeep: %f - SIZE DE MORT : %d, PQ SIZE : %d\n", current.g, len(puzMap), puzQueue.Len())
		// fmt.Println(current.puzMap)
		// fmt.Println()
		// * on cree les 4(max) chemin depuis current (le top du top)
		for _, neighbor := range current.findNeighbor(env) {
			// ? check des voisins
			// fmt.Printf("label: %s\n", neighbor.label)
			// fmt.Println(neighbor.puzMap)

			// * suite si ton fils est null
			if neighbor.puzMap == nil {
				continue
			}

			neighborNode := puzMap.get(neighbor)
			// fmt.Printf("label: %t, %t\n", neighborNode.open, neighborNode.close)
			if !neighborNode.open && !neighborNode.close {
				neighborNode.open = true
				heap.Push(puzQueue, neighborNode)
			} else {
				// * c'est just que ce code est pas vraiment utile :p
				if current.f >= neighbor.f {
					if neighborNode.close {
						heap.Push(puzQueue, neighborNode)
						neighborNode.open = true
						neighborNode.close = false
					}
					puzMap[neighbor.label] = neighborNode
				}
			}
		}
		// if k < 4 {
		// 	k++
		// } else {
		// 	return nil
		// }

		// ? il est important de prendre son temps
		// time.Sleep(2 * time.Second)
	}
}
