package main

import "fmt"

func main() {
	name := "arkar"

	fmt.Println("My name is", name)

	age := 12

	fmt.Println(age)

	var numOne int8 = 25

	fmt.Println(numOne)

	fmt.Printf("My name is %v and My age is %v", name, age)

	// array
	// var ages [3]int = [3]int{12, 22, 22}
	var ages = [3]int{12, 22, 22}

	fmt.Println(ages, len(ages))

	names := [3]string{"arkar", "arkar1", "arkar3"}
	fmt.Println(names, len(names))

	//slices
	scores := []int{100, 222, 222}
	fmt.Println("these are the scores", scores)

	scores[2] = 2213213
	fmt.Println("these are the scores", scores)

	newScores := append(scores, 82)
	fmt.Println("these are the scores", newScores)
}
