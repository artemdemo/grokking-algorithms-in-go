package quicksort

import (
    "fmt"
    "testing"
)

func arrToString(arr []int) string {
    return fmt.Sprint(arr)
}

func Test_Quicksort(t *testing.T) {
    t.Run("sortSimpleArray", func(t *testing.T) {
        arr := []int{2}
        resultStr := arrToString(Quicksort(arr))
        expectedResult := "[2]"
        if resultStr != expectedResult {
            t.Fatalf("Quicksort() should return `%v`, got \"%v\" instead", expectedResult, resultStr)
        }
    })

    t.Run("sortArray", func(t *testing.T) {
        arr := []int{2,5,15,1,4}
        resultStr := arrToString(Quicksort(arr))
        expectedResult := "[1 2 4 5 15]"
        if resultStr != expectedResult {
            t.Fatalf("Quicksort() should return `%v`, got \"%v\" instead", expectedResult, resultStr)
        }
    })
}
