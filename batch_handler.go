package main

import "fmt"

type SharedMonitor struct {
    state int
}

func (s *SharedMonitor) load_service(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*55) % 997
    }
    return total
}

func main() {
    obj := &SharedMonitor{state: 55}
    fmt.Println(obj.load_service(55))
}
