package botomat

import (
    "fmt"
    "sync"
    "time"
)

type model int

// Robot model (type). Each robot can work on a different number of tasks at a time.
const (
    _            model = iota
    UNIPEDAL           // one task at a time
    BIPEDAL            // two tasks at a time
    RADIAL             // three tasks at a time
    QUADRUPEDAL        // four tasks at a time
    ARACHNID           // ten tasks at the same time
    AERONAUTICAL       // twenty tasks at the same time
)

// The robot type should be private. Only a factory struct should return an instance of robot.
type robot struct {
    // Model and Name can be changed from the outside.
    Model model
    Name  string

    tasks *sync.Map // Passed down from BotoMat and shared across multiple robots.
    wg    sync.WaitGroup
}

// Simulates work by sleeping.
func (robot *robot) completeTask(key Task) {
    description := key.description
    eta := key.eta
    fmt.Printf("%s working on task %q...\n", robot.Name, description)
    time.Sleep(eta * time.Millisecond)
    fmt.Printf("Task %q - done.\n", description)
}

// Perform work based on the robot model.
func (robot *robot) Work() {
    switch robot.Model {
    case UNIPEDAL:
        robot.workWithLimit(1)
    case BIPEDAL:
        robot.workWithLimit(2)
    case RADIAL:
        robot.workWithLimit(3)
    case QUADRUPEDAL:
        robot.workWithLimit(4)
    case ARACHNID:
        robot.workWithLimit(10)
    case AERONAUTICAL:
        robot.workWithLimit(20)
    default:
        fmt.Println("Model does not exist.")
    }
    robot.wg.Wait()
}

// Limits the maximum number of goroutines a robot can create at a time.
func (robot *robot) workWithLimit(maxGoroutines int) {
    limit := make(chan struct{}, maxGoroutines)
    robot.tasks.Range(func(key, value interface{}) bool {
        // Preventing multiple robots from processing the same task.
        if value == false {
            // Mark the task as taken.
            robot.tasks.Store(key, true)

            limit <- struct{}{}
            robot.wg.Add(1)

            go func(t Task) {
                robot.completeTask(t)
                robot.tasks.Delete(key)

                robot.wg.Done()
                <-limit
            }(key.(Task))
        }
        return true
    })
}
