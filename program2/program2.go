/*
* Authors - Brijesh Patel, Ajay Shankar
 */

package main

import (
	"fmt"
)

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

	var ans string
	fmt.Println("Would you like to delete a group member? (Y/N)")
	fmt.Scan(&ans)

	if ans == "Y" || ans == "y" {
		var name string
		fmt.Println("What is the name of the group member you want to delete?")
		fmt.Scan(&name)

		// Find the index of the student with the given name
		index := -1
		for i := 0; i < len(members); i++ {
			if members[i].name == name {
				index = i
				break
			}
		}

		if index == -1 {
			fmt.Println("There is no student with that name.")
		} else {
			// Remove the student from the slice
			members = append(members[:index], members[index+1:]...)
			fmt.Println("The student", name, "has been removed")
		}
	}

	fmt.Println("These are the names and IDs of the remaining members:")
	for i := 0; i < len(members); i++ {
		fmt.Println(i+1, "-", members[i].name, "(ID:", members[i].id, ")")
	}
}
