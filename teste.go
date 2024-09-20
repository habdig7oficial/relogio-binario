package main

import (
    "time"
)

func main() {
    count := 0
    for {
        println(count, ": Hello, World")
        time.Sleep(time.Millisecond * 1000)
        count++
    }
}
