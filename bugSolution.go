```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int)
        done := make(chan struct{})

        wg.Add(1)
        go func() {
                defer wg.Done()
                for i := 0; i < 10; i++ {
                        ch <- i
                }
                close(done) // Signal completion
        }()

        wg.Add(1)
        go func() {
                defer wg.Done()
                for i := range ch {
                        fmt.Println(i)
                }
                <-done
        }()

        wg.Wait()
}
```