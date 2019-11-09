package binary_search

func BinarySearch(list []int, item int) (int, bool) {
    low := 0
    high := len(list) - 1
    for low <= high {
        mid := (low + high) / 2
        guess := list[mid]
        if guess == item {
            return mid, true
        }
        if guess > item {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return 0, false
}
