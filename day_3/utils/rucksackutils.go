package rucksackutils

func MakePrioMaps() (map[rune]int) {

    prioMap := make(map[rune]int)
    
    x := 'a'
    y := 1

    for y <= 26 {
        prioMap[x] = y
        x += 1
        y += 1
    }

    x = 'A'

    for y <= 52 {
        prioMap[x] = y
        x += 1
        y += 1
    }

    return prioMap
}