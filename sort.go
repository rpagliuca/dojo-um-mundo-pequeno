package main

/* Sort list of distances */
func sort_distances(distances []Distance) []Distance {
    for i := 0; i < len(distances); i++ {
        for j := i + 1; j < len(distances); j++ {
            if distances[i].distance > distances[j].distance {
                tmp := distances[i]
                distances[i] = distances[j]
                distances[j] = tmp
            }
        }
    }
    return distances
}
