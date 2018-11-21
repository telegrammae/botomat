package botomat

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type model int

// Robot model (type). Each robot can work on a different number of tasks at a time.
const (
    _            model = iota
    UNIPEDAL           // one task at a time
    BIPEDAL            // two tasks at a time
    RADIAL             // three task at a time
    QUADRUPEDAL        // four tasks at a time
    ARACHNID           // all tasks at the same time
    AERONAUTICAL       // process a random number of tasks at the same time
)

// The robot type should be private. Only a factory struct should return an instance of robot.
type robot struct {
    // Model and Name can be changed from the outside.
    Model model
    Name  string

    tasks map[string]Task // Passed down from BotoMat and shared across multiple robots.
    mx    sync.Mutex      // For thread-safe deletion of a task.
    wg    sync.WaitGroup
}

// Simulates work by sleeping. Also deletes a task, once it is complete.
func (robot *robot) completeTask(key string) {
    description := robot.tasks[key].description
    eta := robot.tasks[key].eta
    fmt.Printf("Working on task %q...\n", description)
    time.Sleep(eta * time.Millisecond)
    robot.deleteTask(key)
    fmt.Printf("Task %q - done.\n", description)
}

// Delete a task from the shared map of tasks, while preventing concurrent access to map write.
func (robot *robot) deleteTask(key string) {
    robot.mx.Lock()
    defer robot.mx.Unlock()
    delete(robot.tasks, key)
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
        robot.workWithLimit(len(robot.tasks))
    case AERONAUTICAL:
        rand.Seed(time.Now().UTC().UnixNano())
        r := rand.Intn(len(robot.tasks))
        robot.workWithLimit(r)
    default:
        fmt.Println("Model does not exist.")
    }
    robot.wg.Wait()
}

// Limits the maximum number of goroutines a robot can create at a time.
func (robot *robot) workWithLimit(maxGoroutines int) {
    limit := make(chan struct{}, maxGoroutines)
    for k, _ := range robot.tasks {
        limit <- struct{}{}
        robot.wg.Add(1)
        go func(key string) {
            robot.completeTask(key)
            robot.wg.Done()
            <-limit
        }(k)
    }
}
