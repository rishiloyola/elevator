package main_test

import (
	"container/heap"
	"testing"

	elevator "github.com/rishiloyola/elevator"
	"github.com/stretchr/testify/assert"
)

const (
	UP     = 1
	DOWN   = -1
	STEADY = 0
)

func TestFirstEntry(t *testing.T) {
	pickUpQueue := make(elevator.PriorityQueue, 0)
	heap.Init(&pickUpQueue)
	goalQueue := make(elevator.PriorityQueue, 0)
	heap.Init(&goalQueue)
	e := elevator.Elevator{
		Direction:    STEADY,
		CurrentFloor: 0,
		NumFloors:    10,
		PickUpQueue:  pickUpQueue,
		GoalQueue:    goalQueue,
		Status:       "working",
	}
	e.AddPickUpFloor(2, UP, 5) // The elevator is at 0th level and now we are making a call to transfer from 2 to 5th floor
	assert.Equal(t, 5, e.PickUpQueue[0].Destination)
	e.Step()
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 1, e.GoalQueue.Len())
	e.Step()
	assert.Equal(t, 0, e.GoalQueue.Len())
	assert.Equal(t, UP, e.Direction)
	assert.Equal(t, 5, e.CurrentFloor)
	e.Step()
	assert.Equal(t, 0, e.GoalQueue.Len())
	assert.Equal(t, STEADY, e.Direction)
	assert.Equal(t, 5, e.CurrentFloor)
}

func TestTwoEntry(t *testing.T) {
	pickUpQueue := make(elevator.PriorityQueue, 0)
	heap.Init(&pickUpQueue)
	goalQueue := make(elevator.PriorityQueue, 0)
	heap.Init(&goalQueue)
	e := elevator.Elevator{
		Direction:    STEADY,
		CurrentFloor: 0,
		NumFloors:    10,
		PickUpQueue:  pickUpQueue,
		GoalQueue:    goalQueue,
		Status:       "working",
	}
	e.AddPickUpFloor(2, UP, 5) // The elevator is at 0th level and now we are making a call to transfer from 2 to 5th floor
	e.AddPickUpFloor(3, UP, 7)
	assert.Equal(t, 2, e.PickUpQueue.Len())
	e.Step()
	assert.Equal(t, 2, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.CurrentFloor)
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 1, e.PickUpQueue.Len())
	assert.Equal(t, 1, e.GoalQueue.Len())
	assert.Equal(t, 2, e.CurrentFloor)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 2, e.GoalQueue.Len())
	assert.Equal(t, 3, e.CurrentFloor)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 1, e.GoalQueue.Len())
	assert.Equal(t, 5, e.CurrentFloor)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.GoalQueue.Len())
	assert.Equal(t, 7, e.CurrentFloor)
	e.Step()
	assert.Equal(t, STEADY, e.Direction)
}

func TestUPandDown(t *testing.T) {
	pickUpQueue := make(elevator.PriorityQueue, 0)
	heap.Init(&pickUpQueue)
	goalQueue := make(elevator.PriorityQueue, 0)
	heap.Init(&goalQueue)
	e := elevator.Elevator{
		Direction:    0,
		CurrentFloor: 0,
		NumFloors:    10,
		PickUpQueue:  pickUpQueue,
		GoalQueue:    goalQueue,
		Status:       "working",
	}
	e.AddPickUpFloor(2, UP, 5) // The elevator is at 0th level and now we are making a call to transfer from 2 to 5th floor
	e.AddPickUpFloor(7, DOWN, 2)
	assert.Equal(t, 2, e.PickUpQueue.Len())
	e.Step()
	assert.Equal(t, 2, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.CurrentFloor)
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 1, e.PickUpQueue.Len())
	assert.Equal(t, 1, e.GoalQueue.Len())
	assert.Equal(t, 2, e.CurrentFloor)
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 1, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.GoalQueue.Len())
	assert.Equal(t, 5, e.CurrentFloor)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 1, e.GoalQueue.Len())
	assert.Equal(t, 7, e.CurrentFloor)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 1, e.GoalQueue.Len())
	assert.Equal(t, 7, e.CurrentFloor)
	assert.Equal(t, DOWN, e.Direction)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.GoalQueue.Len())
	assert.Equal(t, 2, e.CurrentFloor)
	assert.Equal(t, DOWN, e.Direction)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.GoalQueue.Len())
	assert.Equal(t, 2, e.CurrentFloor)
	assert.Equal(t, STEADY, e.Direction)
}

func TestSameFloor(t *testing.T) {
	pickUpQueue := make(elevator.PriorityQueue, 0)
	heap.Init(&pickUpQueue)
	goalQueue := make(elevator.PriorityQueue, 0)
	heap.Init(&goalQueue)
	e := elevator.Elevator{
		Direction:    0,
		CurrentFloor: 0,
		NumFloors:    10,
		PickUpQueue:  pickUpQueue,
		GoalQueue:    goalQueue,
		Status:       "working",
	}
	e.AddPickUpFloor(2, UP, 5) // The elevator is at 0th level and now we are making a call to transfer from 2 to 5th floor
	e.AddPickUpFloor(2, UP, 5)
	assert.Equal(t, 2, e.PickUpQueue.Len())
	e.Step()
	assert.Equal(t, 2, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.CurrentFloor)
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 1, e.PickUpQueue.Len())
	assert.Equal(t, 1, e.GoalQueue.Len())
	assert.Equal(t, 2, e.CurrentFloor)
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 2, e.GoalQueue.Len())
	assert.Equal(t, 2, e.CurrentFloor)
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 1, e.GoalQueue.Len())
	assert.Equal(t, 5, e.CurrentFloor)
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.GoalQueue.Len())
	assert.Equal(t, 5, e.CurrentFloor)
	assert.Equal(t, UP, e.Direction)
	e.Step()
	assert.Equal(t, 0, e.PickUpQueue.Len())
	assert.Equal(t, 0, e.GoalQueue.Len())
	assert.Equal(t, 5, e.CurrentFloor)
	assert.Equal(t, STEADY, e.Direction) // It will take one more step to reach to steady state
}
