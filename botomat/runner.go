package botomat

import (
    "fmt"
    "log"
)

func Run() {
    fmt.Println("Plese enter the robot's name:")
    var name string
    _, err := fmt.Scanf("%s", &name)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Plese enter the robot's model name as an integer from 1 to 6:")
    var modelNumber int
    _, err = fmt.Scanf("%d", &modelNumber)
    if err != nil {
        log.Fatal(err)
    }

    m := GenerateRandomTasks(50)

    factory := BotoMat{m}
    robot := factory.NewRobot(model(modelNumber), name)
    go robot.Work()

    anotherRobot := factory.NewRobot(model(RADIAL), "DOS")
    go anotherRobot.Work()

    yetAnotherRobot := factory.NewRobot(model(RADIAL), "TRES")
    yetAnotherRobot.Work()

    fmt.Println("All robots done.")
}
