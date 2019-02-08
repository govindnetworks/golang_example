package main 

import (
	"fmt"
	"path"
)

func main() {
	paths := [] string {
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}
	for _, p := range paths {
		fmt.Println(path.Clean(p))
	}
	fmt.Println(path.Base("/a/b/c/"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
	fmt.Println(path.Ext("/a/b/c/test.css"))
	fmt.Println(path.Split("/a/b/c/d/myfile.css"))
	fmt.Println(path.Join("a", "b","c"))
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Dir("a/b/c"))
	fmt.Println(path.Dir("/a/"))
	fmt.Println(path.Dir("a/"))
	fmt.Println(path.Dir("/"))
	fmt.Println(path.Dir(""))
}
