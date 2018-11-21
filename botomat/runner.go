package botomat

import (
    "fmt"
    "log"
)

func Runner() {
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

    // m := make(map[string]Task)
    // for _, task := range tasks {
    //     key := GetRandomString(8)
    //     m[key] = task
    // }

    m := GenerateRandomTasks(50)

    factory := BotoMat{m}
    robot := factory.NewRobot(model(modelNumber), name)
    robot.Work()

    fmt.Println("All robots done.")
}

var tasks = []Task{
    Task{
        description: "do the dishes",
        eta:         1000,
    }, Task{
        description: "sweep the house",
        eta:         3000,
    }, Task{
        description: "do the laundry",
        eta:         10000,
    }, Task{
        description: "take out the recycling",
        eta:         4000,
    }, Task{
        description: "make a sammich",
        eta:         7000,
    }, Task{
        description: "mow the lawn",
        eta:         20000,
    }, Task{
        description: "rake the leaves",
        eta:         18000,
    }, Task{
        description: "give the dog a bath",
        eta:         14500,
    }, Task{
        description: "bake some cookies",
        eta:         8000,
    }, Task{
        description: "wash the car",
        eta:         20000,
    },
}
