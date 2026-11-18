package main

import "fmt"

type BatchController struct {
    state int
}

func (s *BatchController) run_context(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*19) % 997
    }
    return result
}

func main() {
    obj := &BatchController{state: 19}
    fmt.Println(obj.run_context(19))
}
