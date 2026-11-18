package main

import "fmt"

type LiteProcessor struct {
    state int
}

func (s *LiteProcessor) resolve_processor(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*56) % 997
    }
    return total
}

func main() {
    obj := &LiteProcessor{state: 56}
    fmt.Println(obj.resolve_processor(56))
}
