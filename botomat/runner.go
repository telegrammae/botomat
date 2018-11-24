package botomat

import (
    "flag"
    "fmt"
    "log"
)

var robotCount = flag.Int("robots", 1, "how many robots to run at the same time")

// Exported function to run the program.
func Run() {
    flag.Parse()
    var name string
    var modelNumber int
    var robots = []*robot{}

    m := GenerateRandomTasks(50)
    factory := BotoMat{m}

    for i := 0; i < *robotCount; i++ {
        fmt.Println("Plese enter the robot's name:")
        _, err := fmt.Scanf("%s", &name)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Plese enter the robot's model name as an integer from 1 to 6:")
        fmt.Printf("%d: %s\t - robot performs one task at a time.\n", model(UNIPEDAL), "UNIPEDAL")
        fmt.Printf("%d: %s\t - robot performs two tasks at a time.\n", model(BIPEDAL), "BIPEDAL")
        fmt.Printf("%d: %s\t - robot performs three tasks at a time.\n", model(RADIAL), "RADIAL")
        fmt.Printf("%d: %s\t - robot performs four tasks at a time.\n", model(QUADRUPEDAL), "QUADRUPEDAL")
        fmt.Printf("%d: %s\t - robot performs ten tasks at a time.\n", model(ARACHNID), "ARACHNID")
        fmt.Printf("%d: %s\t - robot performs twenty tasks at a time.\n", model(AERONAUTICAL), "AERONAUTICAL")

        _, err = fmt.Scanf("%d", &modelNumber)
        if err != nil {
            log.Fatal(err)
        }

        robots = append(robots, factory.NewRobot(model(modelNumber), name))
    }

    // The design is such that each additional robot becomes the last one for which the main goroutine waits. Each previous robot starts work in its own goroutine. The last is launched in the main goroutine.
    for _, r := range robots[:len(robots)-1] {
        go r.Work()
    }
    robots[len(robots)-1].Work()

    fmt.Println("All robots done.")
}
