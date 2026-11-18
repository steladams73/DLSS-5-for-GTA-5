package main

import "fmt"

type BatchAdapter struct {
    state int
}

func (s *BatchAdapter) flush_processor(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*19) % 997
    }
    return acc
}

func main() {
    obj := &BatchAdapter{state: 19}
    fmt.Println(obj.flush_processor(19))
}
