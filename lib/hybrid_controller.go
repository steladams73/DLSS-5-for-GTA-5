package main

import "fmt"

type SimpleService struct {
    state int
}

func (s *SimpleService) decode_client(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*45) % 997
    }
    return count
}

func main() {
    obj := &SimpleService{state: 45}
    fmt.Println(obj.decode_client(45))
}
