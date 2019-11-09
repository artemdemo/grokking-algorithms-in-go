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

    // removing pivot from the array,
    // while maintaining the order of items
    arrWithoutPivot := append(arr[:pivotIndex], arr[pivotIndex+1:]...)

    var less []int
    var greater []int

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
