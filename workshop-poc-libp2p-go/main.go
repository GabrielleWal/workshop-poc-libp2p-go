package main

import (
    "fmt"
    "ws/node"
)

func main() {
    myNode := node.NewBootNode()

    err := myNode.Start()
    if err != nil {
        return
    }
    fmt.Println(myNode.GetMultiAddress())
    myNode.Stop()
}