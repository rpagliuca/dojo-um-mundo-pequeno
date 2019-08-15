package main

/* Latitude and longitude data for every friend */
func get_data() map[string]Position {
    friends := map[string]Position {
        "Amigo1": {0, 0},
        "Amigo2": {101, 101},
        "Amigo3": {0, 2},
        "Amigo4": {0, 3},
        "Amigo5": {0, 4},
        "Amigo6": {100, 100},
        "Amigo7": {100, 101},
        "Amigo8": {0, 7},
        "Amigo9": {0, 8},
        "Amigo10": {0, 9},
    }
    return friends
}
