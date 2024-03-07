// structures allow data to be stored in groups
// similar to "class" in other programming languages
// each data point in the structure is called a field

package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	// instantiating structure
	user1 := User{"john", 12}
	user2 := User{
		name: "max",
		age:  17,
	}
	user3 := User{}

	fmt.Println(user1, user2, user3)

	// accessing fields
	user3.name = "donald"
	user3.age = 10

	fmt.Println(user3)

	// anonymous structures
	form := struct {
		username, password string
	}{
		username: "valentino",
		password: "rossi",
	}

	fmt.Println(form)
}
