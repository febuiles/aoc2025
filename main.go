package main

import (
    "bufio"
    "os"
	"fmt"
	"strings"
	"sort"
    "strconv"
)

func main() {
	fmt.Println("main")
    file, err := os.Open("input"    )
    if err != nil {
        panic(err)
    }
    defer file.Close()
    list_a := []int{}
    list_b := []int{}
    distance := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        values := strings.Split(line, "   ")
        first, _ := strconv.Atoi(values[0])
        second, _ := strconv.Atoi(values[1])
        list_a = append(list_a, first)
        list_b = append(list_b, second)
        
    }
    sort.Ints(list_a)
    sort.Ints(list_b)
    for i := 0; i < len(list_a); i++ {
        distance += list_a[i] * count(list_b, list_a[i])
    }
    fmt.Printf("Distance: %d\n", distance)
}

func count(list []int, target int) int {
    counts := make(map[int]int)
    for _, num := range list {
        counts[num]++
    }
    return counts[target]
}


