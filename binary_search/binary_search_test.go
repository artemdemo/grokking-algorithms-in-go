package binary_search

import "testing"

func Test_BinarySearch(t *testing.T) {
    t.Run("itemNotIncluded", func(t *testing.T) {
        list := []int{1,2,3}
        result, ok := BinarySearch(list, 4)

        if ok == true {
            t.Fatalf("BinarySearch() should return `ok=false`, got \"%v\" instead", ok)
        }
        if result != 0 {
            t.Fatalf("BinarySearch() should return `result=0`, got \"%v\" instead", result)
        }
    })

    t.Run("itemIncludedOnTheRight", func(t *testing.T) {
        list := []int{1,2,3,4,5,6,7,9}
        result, ok := BinarySearch(list, 7)

        if ok == false {
            t.Fatalf("BinarySearch() should return `ok=false`, got \"%v\" instead", ok)
        }
        if result != 6 {
            t.Fatalf("BinarySearch() should return `result=6`, got \"%v\" instead", result)
        }
    })

    t.Run("itemIncludedOnTheLeft", func(t *testing.T) {
        t.Name()
        list := []int{1,2,3,4,5,6,7,9}
        result, ok := BinarySearch(list, 2)

        if ok == false {
            t.Fatalf("BinarySearch() should return `ok=false`, got \"%v\" instead", ok)
        }
        if result != 1 {
            t.Fatalf("BinarySearch() should return `result=1`, got \"%v\" instead", result)
        }
    })
}
