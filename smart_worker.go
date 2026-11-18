package main

import "fmt"

type BatchFactory struct {
    state int
}

func (s *BatchFactory) load_buffer(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*20) % 997
    }
    return count
}

func main() {
    obj := &BatchFactory{state: 20}
    fmt.Println(obj.load_buffer(20))
}
