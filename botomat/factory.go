package botomat

import (
    "math/rand"
    "sync"
    "time"
)

// Factory struct that returns instances of robots and keeps a list of tasks.
type BotoMat struct {
    Tasks *sync.Map
}

func (b *BotoMat) NewRobot(model model, name string) *robot {
    return &robot{Model: model, Name: name, wg: sync.WaitGroup{}, tasks: b.Tasks}
}

func GenerateRandomTasks(size int) *sync.Map {
    rand.Seed(time.Now().UTC().UnixNano())
    result := sync.Map{}
    verbs := []string{"walk", "help", "clean", "start", "remember", "make", "open", "close", "move",
        "collect", "get", "clear", "remove", "add", "call", "create", "watch", "fix", "paint", "buy",
        "research", "find", "discover", "welcome", "invite", "freeze", "burn", "build",
        "explore", "water", "play", "study", "share", "donate", "sell", "throw", "hide"}
    nouns := []string{"dog", "cat", "leaves", "kitchen", "floor", "house", "room", "door", "car",
        "garden", "window", "computer", "ceiling", "book", "teapot", "cup", "pen", "website",
        "language", "album", "ticket", "fridge", "lamp", "bell", "fork", "water", "space",
        "key", "bottle", "tire", "clay", "store", "basket", "spoon", "phone", "game"}

    for i := 0; i < size; i++ {
        verb := verbs[rand.Intn(len(verbs)-1)]
        noun := nouns[rand.Intn(len(nouns)-1)]
        description := verb + " the " + noun
        eta := (rand.Intn(5) * 1000) + 1
        result.Store(Task{description, time.Duration(eta)}, false)
    }

    return &result
}

type Task struct {
    description string
    eta         time.Duration
}
