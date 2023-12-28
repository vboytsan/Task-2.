package main

import (
    "fmt"
    "sort"
    "testing"
    "github.com/stretchr/testify/assert"
)
func findMinMaxSums(arr []int) (minSum, maxSum int) {
    if len(arr) < 5 {
        return 0, 0
    }

    sort.Ints(arr)

    minSum = arr[0] + arr[1] + arr[2] + arr[3]

    maxSum = arr[1] + arr[2] + arr[3] + arr[4]

    return minSum, maxSum
}

func TestFindMinMaxSums(t *testing.T) {

    arr := []int{1, 2, 3, 4, 5}
    minSum, maxSum := findMinMaxSums(arr)
    assert.Equal(t, 10, minSum)
    assert.Equal(t, 14, maxSum)

    arr = []int{5, 5, 5, 5, 5}
    minSum, maxSum = findMinMaxSums(arr)
    assert.Equal(t, 20, minSum)
    assert.Equal(t, 25, maxSum)
}

func main() {

    arr := []int{1, 2, 3, 4, 5}

    minSum, maxSum := findMinMaxSums(arr)
    fmt.Printf("Output: %d %d\n", minSum, maxSum)
}

