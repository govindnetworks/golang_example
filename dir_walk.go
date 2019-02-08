package main 

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkhelper(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("prevent panic by handling failure %q: %v",path,err)
	}
	if info.IsDir() && info.Name() == "skip" {
		fmt.Printf("skipping a dir without errors ", info.Name())
		return filepath.SkipDir
	}
	fmt.Println("Dir info ", info.Name())
	return nil
}

func CreateTestDirTree(dirpath string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", fmt.Errorf("Error creating temp dir: %v\n ", err)
	}
	err = os.MkdirAll(filepath.Join(tmpDir, dirpath), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}
	return tmpDir, nil
}

func main() {
	tempDir , err := CreateTestDirTree("test/sample/golang/example")
	if err != nil {
		fmt.Printf("unable to create dir %v", err)
		return
	}
	defer os.RemoveAll(tempDir)
	os.Chdir(tempDir)
	fmt.Println("director to walk ",tempDir)
	fmt.Println("on Unix")
	err = filepath.Walk(".", walkhelper)
	if err != nil {
		fmt.Printf("Error walking path %q %v ", tempDir,err)
	} 
}
