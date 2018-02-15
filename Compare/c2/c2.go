package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	// "log"
)

var (
	sourcefileList, targetfileList []string
	sourceFileCheck, targetFileCheck []string
)

func modified(target, source string) {
	f1, err1 := ioutil.ReadFile(target)
	if err1 != nil {
		// log.Fatal(err1)
	}
	f2, err2 := ioutil.ReadFile(source)
	if err2 != nil {
		// log.Fatal(err2)
	}
	if !bytes.Equal(f1, f2) {
		fmt.Printf("%s MODIFIED\n", target)
	}
}

func deleted() {
	a := 1
	for i := 1; i < len(targetfileList); i++ {
		for j := 1; j < len(targetfileList); j++ {
			if targetfileList[i] == sourcefileList[j] {
				a = 0
				modified(targetFileCheck[i], sourceFileCheck[j])
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
}

func new() {
	a:=1
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

func main() {
	targetDir := "./target"
	sourceDir := "./source"

	terr := filepath.Walk(targetDir, func(path string, f os.FileInfo, err error) error {
		targetFileCheck = append(targetFileCheck, path)
		path = strings.TrimPrefix(path, "target/")
		targetfileList = append(targetfileList, path)
		return nil
	})
	if terr != nil {
		fmt.Println("target ", terr)
	}

	serr := filepath.Walk(sourceDir, func(path string, f os.FileInfo, err error) error {
		sourceFileCheck = append(sourceFileCheck, path)
		path = strings.TrimPrefix(path, "source/")
		sourcefileList = append(sourcefileList, path)
		return nil
	})
	if serr != nil {
		fmt.Println("source ", serr)
	}

	deleted()
	new()

}
