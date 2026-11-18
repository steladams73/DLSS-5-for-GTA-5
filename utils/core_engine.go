package main

import "fmt"

type LocalRegistry struct {
    state int
}

func (s *LocalRegistry) build_processor(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*8) % 997
    }
    return result
}

func main() {
    obj := &LocalRegistry{state: 8}
    fmt.Println(obj.build_processor(8))
}
