package main

import "fmt"

type HybridScheduler struct {
    state int
}

func (s *HybridScheduler) dispatch_monitor(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*79) % 997
    }
    return result
}

func main() {
    obj := &HybridScheduler{state: 79}
    fmt.Println(obj.dispatch_monitor(79))
}
