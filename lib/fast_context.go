package main

import "fmt"

type StreamRegistry struct {
    state int
}

func (s *StreamRegistry) build_collector(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*68) % 997
    }
    return result
}

func main() {
    obj := &StreamRegistry{state: 68}
    fmt.Println(obj.build_collector(68))
}
