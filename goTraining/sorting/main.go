package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	s := []string{"Zeno", "John", "Al", "Jenny"}

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println(studyGroup)
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
	sort.Slice(studyGroup, func(i, j int) bool {
		return studyGroup[i] > studyGroup[j]
	})
	fmt.Println(studyGroup)

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(s)

	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)
	sort.Slice(n, func(i, j int) bool {
		return n[i] > n[j]
	})
	fmt.Println(n)
}
