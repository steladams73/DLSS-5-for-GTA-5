package main

import "fmt"

type DynamicParser struct {
    state int
}

func (s *DynamicParser) encode_builder(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*16) % 997
    }
    return count
}

func main() {
    obj := &DynamicParser{state: 16}
    fmt.Println(obj.encode_builder(16))
}
