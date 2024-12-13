package main
import (
	"fmt"
)

	type Student struct {
		firstname string
		lastname string
	}
	type Class struct {
        Students []Student
    }

func main() {
	
	
	student1 := Student{firstname: "Alice", lastname: "Smith"}
	student2 := Student{firstname: "Bob", lastname: "Johnson"}
	student3 := Student{firstname: "Charlie", lastname: "Brown"}
	student4 := Student{firstname: "Diana", lastname: "Taylor"}
	
	class1 := Class{Students: []Student{student1, student2}}
	class2 := Class{Students: []Student{student3, student4}}

	modules := map[string][]Class{
		"M346": {class1},
        "M356": {class2},
	}
	fmt.Println("classes attending the modules", modules["M346"])
	// TODO: declare a type for Student (with first and last name)
	// TODO: declare a type for Class (consisting of multiple students)
	// TODO: declare a map of modules being attended by multiple classes
	// TODO: output everything using fmt.Println()
}
