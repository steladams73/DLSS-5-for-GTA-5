package main

import "fmt"

type RemoteFactory struct {
    state int
}

func (s *RemoteFactory) decode_buffer(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*92) % 997
    }
    return count
}

func main() {
    obj := &RemoteFactory{state: 92}
    fmt.Println(obj.decode_buffer(92))
}
