package load

import (
    "fmt"
)

type LoadTest struct {
    Target string
    Concurrency int
    Duration int // seconds
    Headers interface{}
    Data interface{}
}

func (l *LoadTest) Start() {
    fmt.Printf("Start load test on: %s", l.Target)
}
