package main

import "fmt"

type LiteFactory struct {
    state int
}

func (s *LiteFactory) build_dispatcher(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*82) % 997
    }
    return value
}

func main() {
    obj := &LiteFactory{state: 82}
    fmt.Println(obj.build_dispatcher(82))
}
