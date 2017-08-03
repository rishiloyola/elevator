# elevator

### Problem
Design and implement an elevator control system. What data structures, interfaces and algorithms will you need? Your elevator control system should be able to handle a few elevators -- up to 16.

### Key Functions
* `Status() []State`: Status returns `id` wise current floor at which elevator is situated and GoalFloor at which elevator is heading. For the elevator which is in the steady state will show the `-1` as a GoalFloor.
* `PickUp(floor int, destination int)`: This function helps to schedule the users among the different elevators. User will make a request to `ElevatorControlSystem` by providing the floor from which he wants to enter the elevator and wanted to go to mentioned floor.
* `Step()`: Allows one unit of time to pass.
* `Update(id int, floor, goalFloor int) []Condition`: This function helps to receive an update about the status of an elevator. It will return `working` and `notworking` as an output for each elevator.

### Assumptions
* Elevator does not have any kind of limitations. It can always pick up as much as it can, in the same direction where it is heading.
* It will serve only one user at each and every step. So if N user wants to go to the same floor together, it will require N steps to let them come out of an elevator.

### Scheduling Algorithm
* My scheduling algorithm is based on the elevator algorithm ([SCAN](https://en.wikipedia.org/wiki/Elevator_algorithm)).
* Initially all the elevators are situated at 0th floor in a `steady` state.
* When the requests come in elevators start moving in that direction(Upward).
* At first elevators pickup the users from the requested floor and then transfer them to their respective destinations.
* Elevators move in the same direction as long as they have the pending requests stored in a priority queue for that particular direction.
* When the priority queue is empty the elevator will go into a steady state and change the direction if there are requests in the opposite direction.
* I used two priority queues to differentiate the pickup requests and transfer requests. The time complexity for insertion is O(logN). While time complexity of search is O(N). I can further optimize search complexity and can reduce it to O(logN)
* Users are allowed to predefine their destinations so that we can schedule them in a better way.
* This algorithm is at least better than FCFS because it takes pickup request's direction into considering whether an elevator should pick up the request or not.
* This system can serve multiple requests concurrently by adding more then one elevators.

### Improvements
* [Destination dispatch](https://en.wikipedia.org/wiki/Destination_dispatch) is an optimization technique we can use while having more then one elevator in our system. Basically it divides all the elevators in particular groups that these groups all pickup requests for the same destinations into the same elevator. This algorithm helps to reduce waiting and traveling time.
* Another improvement in algorithm would be to pick the elevator with the least entries in its queues and the closest to the pickup floor or requested destination. For each and every request we can calculate specific score and based on this score we can select the elevator.
* I can add goroutines to make it truly simulate the multi-elevators system.
* If we have a certain pattern for our pickup points then we can further analyze this pattern and can make a **ML model** out of it. Which can simulate the elevators by learning through past data.
* I can make it more robust and could come up with better design pattern if time was not limited up to 4 hours.

### Setup

Setup the `go1.6` version on your machine and follow the following instructions.
```
go get -u github.com/rishiloyola/elevator
go get
go build
./elevator
```
To simulate or to add more cases edit the `main.go` file.

To test the code run following command.
```
go test elevator_test.go
```
