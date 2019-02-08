package main 


import (
	"fmt"
	"sort"
	
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	people := []Person {
		{"Bob", 43},
		{"John", 42},
	} 
	fmt.Println(people)

	sort.Slice(people, func(i,j int)bool {
		return people[i].Age < people[j].Age
		})
	fmt.Println(people)
}
