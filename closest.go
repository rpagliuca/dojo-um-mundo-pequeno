package main

/* Return the 3 closest friends */
func closest_friends(friends map[string]Position, friend string) []Distance {
    pos1 := friends[friend]
    distances := []Distance{}
    for key, pos := range friends {
        if key == friend {
            continue
        }
        distances = append(distances, Distance{key, distance(pos1, pos)})
    }
    distances = sort_distances(distances)
    return distances[0:3]
}
