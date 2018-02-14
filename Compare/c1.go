package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
	targetDir := "./target"
	sourceDir := "./source"

    targetfileList := []string{}
    terr := filepath.Walk(targetDir, func(path string, f os.FileInfo, err error) error {
        targetfileList = append(targetfileList, path)
        return nil
	})
	if terr != nil {
		fmt.Println("target ", terr)
	}

	sourcefileList := []string{}
    serr := filepath.Walk(sourceDir, func(path string, f os.FileInfo, err error) error {
		fmt.Println(path)
        sourcefileList = append(sourcefileList, path)
        return nil
	})
	if serr != nil {
		fmt.Println("source ", serr)
	}
	
    if targetfileList[0] == sourcefileList[0] {
		fmt.Println("ok")
	}

}
