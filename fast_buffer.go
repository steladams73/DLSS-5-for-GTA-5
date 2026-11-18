package main

import "fmt"

type StreamProcessor struct {
    state int
}

func (s *StreamProcessor) resolve_worker(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*25) % 997
    }
    return result
}

func main() {
    obj := &StreamProcessor{state: 25}
    fmt.Println(obj.resolve_worker(25))
}
