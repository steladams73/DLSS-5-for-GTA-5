package main

import "fmt"

type SimpleWorker struct {
    state int
}

func (s *SimpleWorker) decode_cache(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*78) % 997
    }
    return result
}

func main() {
    obj := &SimpleWorker{state: 78}
    fmt.Println(obj.decode_cache(78))
}
