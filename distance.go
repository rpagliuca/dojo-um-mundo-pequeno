package main

import (
    "math"
)

/* Calculate the distance between 2 positions */
func distance(pos1 Position, pos2 Position) float64 {
    return math.Pow(math.Pow(pos1.lat - pos2.lat, 2.0) + math.Pow(pos1.lng - pos2.lng, 2.0), 0.5)
}
