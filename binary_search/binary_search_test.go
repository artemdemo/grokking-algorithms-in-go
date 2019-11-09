package binary_search

import "testing"

func itemNotIncluded(t *testing.T) {
    list := []int{1,2,3}
    result, ok := BinarySearch(list, 4)

    if ok == true {
        t.Fatalf("BinarySearch() should return `ok=false`, got \"%v\" instead", ok)
    }
    if result != 0 {
        t.Fatalf("BinarySearch() should return `result=0`, got \"%v\" instead", result)
    }
}

func itemIncludedOnTheRight(t *testing.T) {
    list := []int{1,2,3,4,5,6,7,9}
    result, ok := BinarySearch(list, 7)

    if ok == false {
        t.Fatalf("BinarySearch() should return `ok=false`, got \"%v\" instead", ok)
    }
    if result != 6 {
        t.Fatalf("BinarySearch() should return `result=6`, got \"%v\" instead", result)
    }
}

func itemIncludedOnTheLeft(t *testing.T) {
    list := []int{1,2,3,4,5,6,7,9}
    result, ok := BinarySearch(list, 2)

    if ok == false {
        t.Fatalf("BinarySearch() should return `ok=false`, got \"%v\" instead", ok)
    }
    if result != 1 {
        t.Fatalf("BinarySearch() should return `result=1`, got \"%v\" instead", result)
    }
}

func Test_BinarySearch(t *testing.T) {
    itemNotIncluded(t)
    itemIncludedOnTheRight(t)
    itemIncludedOnTheLeft(t)
}
