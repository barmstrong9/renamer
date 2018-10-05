package main

import (
	"path/filepath"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
	"fmt"
)

func main() {
	// fileName := "birthday_001.txt"
	// newName, err := match(fileName, 4)
	// if err != nil{
	// 	fmt.Println("No match")
	// 	os.Exit(1)
	// }
	// fmt.Println(newName)
	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil{
		panic(err)
	}
	count := 0
	var toRename []string

	for _, file := range files {
		if file.IsDir(){
		} else {
			_, err := match(file.Name(), 0) 
			if err == nil{
				count++	
				toRename = append(toRename, file.Name())
			}
		}
	}
	for _, origFileName:= range toRename{
		origPath := filepath.Join(dir, origFileName)
		newFileName, err := match(origFileName, count)
		if err != nil{
			panic(err)
		}
		newPath := filepath.Join(dir, newFileName)
		fmt.Printf("mv %s => %s\n", origPath, newPath)
		err = os.Rename(origPath, newPath)
		if err != nil{
			panic(err)
		}
	}
}

func match (filename string, total int) (string, error){
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])

	if err != nil{
		return "", fmt.Errorf("%s did not match our pattern", filename)
	}
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}