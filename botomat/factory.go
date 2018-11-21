package botomat

import (
    "encoding/hex"
    "math/rand"
    "sync"
    "time"
)

// Factory struct that returns instances of robots and keeps a list of tasks.
type BotoMat struct {
    Tasks map[string]Task
}

func (b *BotoMat) NewRobot(model model, name string) *robot {
    return &robot{Model: model, Name: name, wg: sync.WaitGroup{}, tasks: b.Tasks}
}

// Returns n random tasks for a robot from a list of tasks.
// Marks taken tasks.
// func (b *BotoMat) GetTasks(n int) []Task {
//     rand.Seed(time.Now().UTC().UnixNano())
//     // Get a set of n unique random tasks.
//     result := []Task{}
//     for _, taskIndex := range rand.Perm(n) {
//         task := b.Tasks[taskIndex]
//         if !task.taken {
//             result = append(result, task)
//             tasks[taskIndex].taken = true
//         }
//     }
//     return result
// }

func GenerateRandomTasks(size int) map[string]Task {
    rand.Seed(time.Now().UTC().UnixNano())
    result := make(map[string]Task, size)
    verbs := []string{"walk", "help", "clean", "start", "remember", "make", "open", "close", "move", "collect", "get", "clear", "remove", "add", "call", "create", "watch", "fix", "paint", "buy"}
    nouns := []string{"dog", "cat", "leaves", "kitchen", "floor", "house", "room", "door", "car", "garden", "window", "computer", "ceiling", "book"}

    for i := 0; i < size; i++ {
        key := GetRandomString(8)
        verb := verbs[rand.Intn(len(verbs)-1)]
        noun := nouns[rand.Intn(len(nouns)-1)]
        description := verb + " the " + noun
        eta := (rand.Intn(10) * 1000) + 1
        result[key] = Task{description, time.Duration(eta)}
    }

    return result
}

// Gets a random string identifier. The size parameter indicates the desired string length.
func GetRandomString(size int) string {
    b := make([]byte, size/2)
    rand.Read(b)
    return hex.EncodeToString(b)
}

type Task struct {
    description string
    eta         time.Duration
}
