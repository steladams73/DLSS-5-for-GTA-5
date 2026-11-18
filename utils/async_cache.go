package main

import "fmt"

type CoreContext struct {
    state int
}

func (s *CoreContext) resolve_manager(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*17) % 997
    }
    return value
}

func main() {
    obj := &CoreContext{state: 17}
    fmt.Println(obj.resolve_manager(17))
}
