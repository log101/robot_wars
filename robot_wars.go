package main

import (
	"fmt"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 2, 3}

	map1 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	map2 := map[int]string{
		3: "c",
		2: "b",
		1: "a",
	}
	fmt.Print(slices.Compare(arr1, arr2))
	fmt.Print(maps.Equal(map1, map2))

}
