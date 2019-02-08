
package main 

import (
	"fmt"
	"log"
	"io"
	"os"
	"crypto/md5"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hash := md5.New()

	if _,err = io.Copy(hash, f); err !=nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", hash.Sum(nil))

}
