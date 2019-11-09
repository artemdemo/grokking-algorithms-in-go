package quicksort

import (
    "math/rand"
    "time"
)

func Quicksort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    rand.Seed(time.Now().Unix())            // initialize global pseudo random generator
    pivotIndex := rand.Intn(len(arr) - 1)   // selecting random element from a given array
    pivot := arr[pivotIndex]
    var less []int
    var greater []int
    arrWithoutPivot := append(arr[:pivotIndex], arr[pivotIndex+1:]...)
    for _, num := range arrWithoutPivot {
        if num <= pivot {
            less = append(less, num)
        } else {
            greater = append(greater, num)
        }
    }
    return append(
        append(Quicksort(less), pivot),
        Quicksort(greater)...
    )
}
