package botomat

import (
    "testing"
)

func TestTaskListModification(t *testing.T) {
    tasks := map[string]Task{"abc": Task{
        description: "give the dog a bath",
        eta:         500,
    }, "dce": Task{
        description: "bake some cookies",
        eta:         1000,
    }, "fgh": Task{
        description: "wash the car",
        eta:         1200},
    }

    factory := BotoMat{tasks}
    robot := factory.NewRobot(model(BIPEDAL), "r2d2")
    robot.Work()

    if len(tasks) != 0 {
        t.Error("Expected the map of tasks to be emptied but its length is", len(tasks))
    }
}

func TestGenerateRandomTasks(t *testing.T) {
    tasks := GenerateRandomTasks(100)

    if len(tasks) != 100 {
        t.Error("Expected map to contain 100 entries but have", len(tasks))
    }
}

// func TestGetTasks(t *testing.T) {
//     tasks := []Task{Task{
//         description: "give the dog a bath",
//         eta:         14500,
//     }, Task{
//         description: "bake some cookies",
//         eta:         8000,
//     }, Task{
//         description: "wash the car",
//         eta:         20000,
//     }}
//     factory := BotoMat{tasks}

//     subset := factory.GetTasks(2)
//     if len(subset) != 2 {
//         t.Error("Expected 2 tasks, but got", len(subset))
//     }
//     if subset[0] == subset[1] {
//         t.Error("Tasks must be unique.")
//     }
// }
