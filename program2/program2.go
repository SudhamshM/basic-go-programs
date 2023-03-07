package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	var groupSize int
	fmt.Println("Please enter the size of the group:")
	fmt.Scan(&groupSize)

	members := make([]Student, groupSize)

	for i := 0; i < groupSize; i++ {
		var name string
		var id int
		fmt.Println("Please enter the name of the group member:")
		fmt.Scan(&name)
		fmt.Println("Please enter the student ID of the group member:")
		fmt.Scan(&id)
		members[i] = Student{name, id}
	}

	fmt.Println("This group has", groupSize, "members")
	fmt.Println("These are the names and IDs of the members:")

	for i := 0; i < groupSize; i++ {
		fmt.Println(i+1, "-", members[i].name, "(ID:", members[i].id, ")")
	}

	fmt.Println()
}
