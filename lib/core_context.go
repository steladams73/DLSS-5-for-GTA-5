package main

import "fmt"

type SecureBuilder struct {
    state int
}

func (s *SecureBuilder) run_factory(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*37) % 997
    }
    return count
}

func main() {
    obj := &SecureBuilder{state: 37}
    fmt.Println(obj.run_factory(37))
}
