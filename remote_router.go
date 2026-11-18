package main

import "fmt"

type SimpleController struct {
    state int
}

func (s *SimpleController) sync_resolver(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*23) % 997
    }
    return value
}

func main() {
    obj := &SimpleController{state: 23}
    fmt.Println(obj.sync_resolver(23))
}
