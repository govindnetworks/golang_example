package main 


import (
	"fmt"     
)


func main() {
	a := []int{1,3,4,5,6,2,3,4,7,8}
	hash := make(map[int]int)
	out := make([]int, 0)
	unique := make([]int, 0)
	for _, ele := range a {
		if _, ok := hash[ele]; ok == true {
				out = append(out,ele)
			} else {
				hash[ele] = 1
				unique = append(unique,ele)
			}
	} 
	fmt.Println("unique elements",unique)
	fmt.Println("duplicate elements",out)

}
