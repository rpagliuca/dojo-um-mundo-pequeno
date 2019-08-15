package main

import (
    "fmt"
)

/* Main function */
func main() {
    friends := get_data()
    for current_friend, _ := range friends {
        closest := closest_friends(friends, current_friend)
        for i, friend := range closest {
            fmt.Printf("%v - %vº mais próximo: %v (distância: %v)\n", current_friend, i + 1, friend.name, friend.distance)
        }
        fmt.Println("")
    }
}
