// main.go

package main

import (
    "fmt"
)

func main() {
    warState, err := fetchWarState()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Current War State: %+v\n", *warState)
}
