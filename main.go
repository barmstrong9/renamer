package main

import (
	"path/filepath"
	//"io/ioutil"
	"strconv"
	"strings"
	"os"
	"fmt"
)

type file struct{
	name string
	path string
}

func main() {
	dir := "./sample"
	var toRename []file
	filepath.Walk(dir, func(path string, info os.FileInfo, err error)error{
	if info.IsDir(){
		return nil	
	}
	if _, err := match(info.Name()); err == nil {
		toRename = append(toRename, file{
			name: info.Name(),
			path: path,
		})
	}
		return nil
	})
	for _, f := range toRename{
		fmt.Printf("%q\n", f)
	}
	for _, orig:= range toRename{
		var n file
		var err error
		n.name, err =  match(orig.name)
		if err != nil{
			fmt.Println("Error matching: ", orig.path, err.Error())
		}
		n.path = filepath.Join(dir, n.name)

		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		err = os.Rename(orig.path, n.path)
		if err != nil{
			fmt.Println("Error Renaming:", orig.path, err.Error())
		}
	}
}

func match (filename string) (string, error){
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])

	if err != nil{
		return "", fmt.Errorf("%s did not match our pattern", filename)
	}
	return fmt.Sprintf("%s - %d.%s", strings.Title(name), number, ext), nil
}