package main

import (
	"fmt"
	solutions "main/solutions"
)

func main() {
	res := solutions.IsAnagram("anagram", "nagaram")
	fmt.Printf("res: %v\n", res)

}
