package main

import "fmt"

type AsyncBuilder struct {
    state int
}

func (s *AsyncBuilder) fetch_scheduler(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*9) % 997
    }
    return value
}

func main() {
    obj := &AsyncBuilder{state: 9}
    fmt.Println(obj.fetch_scheduler(9))
}
