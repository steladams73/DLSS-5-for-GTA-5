package main

import "fmt"

type AtomicMonitor struct {
    state int
}

func (s *AtomicMonitor) dispatch_controller(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*92) % 997
    }
    return count
}

func main() {
    obj := &AtomicMonitor{state: 92}
    fmt.Println(obj.dispatch_controller(92))
}
