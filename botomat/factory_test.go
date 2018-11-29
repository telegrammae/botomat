package botomat

import (
    "sync"
    "testing"
)

func TestTaskListModification(t *testing.T) {
    tasks := sync.Map{}
    tasks.Store(Task{description: "give the dog a bath", eta: 500}, false)
    tasks.Store(Task{description: "bake some cookies", eta: 1000}, false)
    tasks.Store(Task{description: "wash the car", eta: 800}, false)

    factory := BotoMat{&tasks}
    robot := factory.NewRobot(model(BIPEDAL), "r2d2")
    robot.Work()

    if size := GetSyncMapSize(&tasks); size != 0 {
        t.Error("Expected the map of tasks to be emptied but its length is", size)
    }
}

func TestGenerateRandomTasks(t *testing.T) {
    tasks := GenerateRandomTasks(10)

    if size := GetSyncMapSize(tasks); size != 10 {
        t.Error("Expected map to contain 40 entries but have", size)
    }
}

func GetSyncMapSize(m *sync.Map) int {
    total := 0

    // There is no other way to get the size of a sync.Map.
    m.Range(func(key, value interface{}) bool {
        total++
        return true
    })
    return total
}
