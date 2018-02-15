package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	a := 1
	targetDir := "./target"
	sourceDir := "./source"

	targetfileList := []string{}
	terr := filepath.Walk(targetDir, func(path string, f os.FileInfo, err error) error {
		path = strings.TrimPrefix(path, "target/")
		targetfileList = append(targetfileList, path)
		return nil
	})
	if terr != nil {
		fmt.Println("target ", terr)
	}

	sourcefileList := []string{}
	serr := filepath.Walk(sourceDir, func(path string, f os.FileInfo, err error) error {
		path = strings.TrimPrefix(path, "source/")
		sourcefileList = append(sourcefileList, path)
		return nil
	})
	if serr != nil {
		fmt.Println("source ", serr)
	}

	//CHECKING DELETED FILES
	for i := 1; i < len(targetfileList); i++ {
		for j := 1; j < len(targetfileList); j++ {
			if targetfileList[i] == sourcefileList[j] {
				a = 0
			} else {
				if a != 0 {
					a++
				}
			}
		}
		if a > 0 {
			fmt.Printf("%s DELETED\n", targetfileList[i])
		} else {
			a = 1
		}
	}

	//CHECKING NEW FILES
	for i := 1; i < len(sourcefileList); i++ {
		for j := 1; j < len(sourcefileList); j++ {
			if sourcefileList[i] == targetfileList[j] {
				a = 0
			} else {
				if a != 0 {
					a++
				}
			}
		}
		if a > 0 {
			fmt.Printf("%s NEW\n", sourcefileList[i])
		} else {
			a = 1
		}
	}
}
