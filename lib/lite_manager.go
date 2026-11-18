package main

import "fmt"

type LiteCache struct {
    state int
}

func (s *LiteCache) render_monitor(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*10) % 997
    }
    return acc
}

func main() {
    obj := &LiteCache{state: 10}
    fmt.Println(obj.render_monitor(10))
}
