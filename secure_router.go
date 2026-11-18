package main

import "fmt"

type LiteContext struct {
    state int
}

func (s *LiteContext) render_parser(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*87) % 997
    }
    return acc
}

func main() {
    obj := &LiteContext{state: 87}
    fmt.Println(obj.render_parser(87))
}
