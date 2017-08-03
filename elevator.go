package main

import (
	"container/heap"
	"fmt"
)

type Elevator struct {
	Id           int
	Direction    int // 0: Not movies, 1: moving up, -1: movin down
	CurrentFloor int
	PickUpQueue  PriorityQueue
	GoalQueue    PriorityQueue
	NumFloors    int
	Status       string
}

func (e *Elevator) AddGoalFloor(goalFloor int, direction int) {
	item := &Item{
		Direction:   direction,
		Floor:       goalFloor,
		priority:    goalFloor,
		Destination: -1,
	}
	heap.Push(&e.GoalQueue, item)
}

func (e *Elevator) AddPickUpFloor(floor int, direction int, dest int) {
	if floor >= e.NumFloors || dest >= e.NumFloors || floor < 0 || dest < 0 || floor == dest {
		fmt.Println("Can't add pickup floor")
		return
	}

	item := &Item{
		Direction:   direction,
		priority:    floor,
		Floor:       floor,
		Destination: dest,
	}
	heap.Push(&e.PickUpQueue, item)
}

func (e *Elevator) GetSeparateQueue() (PriorityQueue, PriorityQueue) {
	var (
		upwardQueue   PriorityQueue
		downwardQueue PriorityQueue
	)
	for i := 0; i < e.PickUpQueue.Len(); i++ {
		switch {
		case e.PickUpQueue[i].Floor >= e.CurrentFloor:
			heap.Push(&upwardQueue, e.PickUpQueue[i])
		case e.PickUpQueue[i].Floor < e.CurrentFloor:
			heap.Push(&downwardQueue, e.PickUpQueue[i])
		}
	}

	for i := 0; i < e.GoalQueue.Len(); i++ {
		switch {
		case e.GoalQueue[i].Direction == UP:
			heap.Push(&upwardQueue, e.GoalQueue[i])
		case e.GoalQueue[i].Direction == DOWN:
			heap.Push(&downwardQueue, e.GoalQueue[i])
		}
	}

	return upwardQueue, downwardQueue
}

func (e *Elevator) NextBiggestFloor(upwardQueue PriorityQueue) *Item {
	for i := upwardQueue.Len() - 1; i >= 0; i-- {
		if upwardQueue[i].Floor >= e.CurrentFloor {
			return upwardQueue[i]
		}
	}
	return nil
}

func (e *Elevator) NextSmallestFloor(downwardQueue PriorityQueue) *Item {
	for i := 0; i < downwardQueue.Len(); i++ {
		if downwardQueue[i].Floor <= e.CurrentFloor {
			return downwardQueue[i]
		}
	}
	return nil
}

func (e *Elevator) RemovePickUpFloor() {
	for i := 0; i < e.PickUpQueue.Len(); i++ {
		if e.PickUpQueue[i].Floor == e.CurrentFloor {
			heap.Remove(&e.PickUpQueue, i)
			break
		}
	}
}

func (e *Elevator) RemoveGoalFloor() {
	for i := 0; i < e.GoalQueue.Len(); i++ {
		if e.GoalQueue[i].Floor == e.CurrentFloor {
			heap.Remove(&e.GoalQueue, i)
			break
		}
	}
}

func (e *Elevator) FindGoalFloor() int {
	upwardQueue, downwardQueue := e.GetSeparateQueue()
	var goalFloor int

	if e.Direction == UP && upwardQueue.Len() > 0 {
		goalFloor = e.NextBiggestFloor(upwardQueue).priority
	} else if e.Direction == DOWN && downwardQueue.Len() > 0 {
		goalFloor = e.NextSmallestFloor(downwardQueue).priority
	} else if e.Direction == STEADY {
		goalFloor = -1 // default goalfloor is -1
	}

	return goalFloor
}

func (e *Elevator) Step() {
	if e.Status != "working" {
		return
	}
	upwardQueue, downwardQueue := e.GetSeparateQueue()
	switch {
	case upwardQueue.Len() == 0 && downwardQueue.Len() == 0:
		e.Direction = STEADY
		return
	case upwardQueue.Len() > 0 && e.Direction == STEADY:
		var direction int
		if upwardQueue[0].Floor > e.CurrentFloor {
			direction = UP
		} else {
			direction = DOWN
		}
		e.Direction = direction
	case downwardQueue.Len() > 0 && e.Direction == STEADY:
		var direction int
		if downwardQueue[0].Floor > e.CurrentFloor {
			direction = UP
		} else {
			direction = DOWN
		}
		e.Direction = direction
	case e.Direction == UP && upwardQueue.Len() == 0 && downwardQueue.Len() > 0:
		e.Direction = DOWN
	case e.Direction == UP && upwardQueue.Len() > 0:
		// Using the current floor find next biggest floor. Now check wheather it is a request floor or goalfloor.
		// If it is goalfloor then remove that floor from upwardqueue. Remove that floor from goalqueue. Update currentfloor.
		// If it is requestfloor then update currentfloor. Add that element in goalfloor. Remove that floor from pickupqueue.
		floor := e.NextBiggestFloor(upwardQueue)
		if floor != nil {
			if floor.Destination != -1 {
				e.CurrentFloor = floor.Floor
				e.AddGoalFloor(floor.Destination, floor.Direction) // TODO: Here it is assuming that one person from one floor.
				e.RemovePickUpFloor()
			} else {
				e.CurrentFloor = floor.Floor
				heap.Remove(&upwardQueue, floor.index)
				e.RemoveGoalFloor()
			}
		} else {
			e.Direction = DOWN
		}
	case e.Direction == DOWN && downwardQueue.Len() == 0 && upwardQueue.Len() > 0:
		e.Direction = UP
	case e.Direction == DOWN && downwardQueue.Len() > 0:
		// Using current floor find the next smallest number from downwardqueue. Identify wheather it is a request floor or goalfloor.
		// If it is goalfloor then remove that floor from downwardqueue. Remove that floor from goalqueue. Update currentfloor.
		// If it is requestfloor then update currentfloor. Add that element in goalfloor. Remove that floor from pickupqueue.
		floor := e.NextSmallestFloor(downwardQueue)
		if floor != nil {
			if floor.Destination != -1 {
				e.CurrentFloor = floor.Floor
				e.AddGoalFloor(floor.Destination, floor.Direction)
				e.RemovePickUpFloor()
			} else {
				e.CurrentFloor = floor.Floor
				heap.Remove(&downwardQueue, floor.index)
				e.RemoveGoalFloor()
			}
		} else {
			e.Direction = UP
		}
	}
}
