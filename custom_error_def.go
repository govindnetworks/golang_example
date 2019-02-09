package main

import (
	"fmt"
	"time"
	"runtime"
	"strings"
)


type FeatureDateError struct {
	When time.Time
	What string
	File string
	Line int
}

func (e FeatureDateError) Error() string {
	 
	return fmt.Sprintf("%v: %s: %d: %v", e.When, e.File,e.Line, e.What)
}

func oops() error {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "linux" && strings.Contains(runtime.GOARCH, "64")  {
		_, file, line, _ := runtime.Caller(1)

		return FeatureDateError{
			time.Now(),
			"Exception: year 2038 problem ", 
			file, 
			line,
		}			
	}

	return nil

}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
