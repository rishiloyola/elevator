package main

import (
	"container/heap"
)

const (
	UP     = 1
	DOWN   = -1
	STEADY = 0
)

type ElevatorControlSystem struct {
	Elevators []Elevator
}

type State struct {
	Id           int
	CurrentFloor int
	GoalFloor    int
}

type Condition struct {
	Id     int
	Status string
}

func (e *ElevatorControlSystem) Status() []State {
	output := make([]State, 0)
	for i := 0; i < len(e.Elevators); i++ {
		output = append(output, State{
			Id:           i,
			CurrentFloor: e.Elevators[i].CurrentFloor,
			GoalFloor:    e.Elevators[i].FindGoalFloor(),
		})
	}
	return output
}

func (e *ElevatorControlSystem) PickUp(floor int, destination int) {
	var direction int

	if destination > floor {
		direction = UP
	} else {
		direction = DOWN
	}
	i := 0
	for ; i < len(e.Elevators); i++ {
		if e.Elevators[i].Direction == STEADY && e.Elevators[i].Status == "working" {
			e.Elevators[i].AddPickUpFloor(floor, direction, destination)
			break
		}
	}
	if i == len(e.Elevators) {
		for ; i < len(e.Elevators); i++ {
			if e.Elevators[i].Status == "working" {
				e.Elevators[i].AddPickUpFloor(floor, direction, destination)
				break
			}
		}
	}
}

func (e *ElevatorControlSystem) Step() {
	for i := 0; i < len(e.Elevators); i++ {
		e.Elevators[i].Step()
	}
}

func (e *ElevatorControlSystem) Update(id int, floor, goalFloor int) []Condition {
	output := make([]Condition, 0)
	for i := 0; i < len(e.Elevators); i++ {
		output = append(output, Condition{
			Id:     i,
			Status: e.Elevators[i].Status,
		})
	}
	return output
}

func NewElevatorController(totalElevator int, totalFloor int) *ElevatorControlSystem {
	controller := &ElevatorControlSystem{}
	pickUpQueue := make(PriorityQueue, 0)
	heap.Init(&pickUpQueue)
	goalQueue := make(PriorityQueue, 0)
	heap.Init(&goalQueue)

	for i := 0; i < totalElevator; i++ {
		controller.Elevators = append(controller.Elevators, Elevator{
			Direction:    STEADY,
			CurrentFloor: 0,
			NumFloors:    totalFloor,
			PickUpQueue:  pickUpQueue,
			GoalQueue:    goalQueue,
			Status:       "working",
		})
	}
	return controller
}
