### Extend Struct

In Go, you cannot directly extend a struct like you would in object-oriented languages such as Python or Java. However, you can achieve similar functionality by embedding one struct within another. 

```bash
package main

import (
    "fmt"
)

// Base struct
type Person struct {
    FirstName string
    LastName  string
}

// Extended struct
type Employee struct {
    Person          // Embedding Person struct
    EmployeeID int
    Department string
}

func main() {
    // Creating an instance of the extended struct
    emp := Employee{
        Person: Person{
            FirstName: "John",
            LastName:  "Doe",
        },
        EmployeeID: 1234,
        Department: "Engineering",
    }

    // Accessing fields from the embedded struct
    fmt.Println("First Name:", emp.FirstName)
    fmt.Println("Last Name:", emp.LastName)

    // Accessing fields from the extended struct
    fmt.Println("Employee ID:", emp.EmployeeID)
    fmt.Println("Department:", emp.Department)
}
```