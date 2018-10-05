package main

import (
	"flag"
	"sort"
	"path/filepath"
	//"io/ioutil"
	"strconv"
	"strings"
	"os"
	"fmt"
)

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", true, "whether or not this is a dry run")
	flag.Parse()
	walkDir := "./sample"
	toRename := make(map[string][]string)
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error)error{
	if info.IsDir(){
		return nil	
	}
	curDir := filepath.Dir(path)
	if m, err := match(info.Name()); err == nil {
		key := filepath.Join(curDir, fmt.Sprintf("%s.%s", m.base, m.ext))
		toRename[key] = append(toRename[key], info.Name())
	}
		return nil
	})

	for key, files := range toRename{
		dir := filepath.Dir(key)
		n := len(files)
		sort.Strings(files)
		for fi, filename := range files{
			res, _ := match(filename)
			newFileName := fmt.Sprintf("%s - %d of %d.%s", res.base, (fi + 1), n, res.ext)
			oldPath := filepath.Join(dir, filename)
			newPath := filepath.Join(dir, newFileName)
			fmt.Printf("mv %s => %s\n", oldPath, newPath)
			if !dry {
				err := os.Rename(oldPath, newPath)
				if err != nil{
					fmt.Println("Error Renaming:", oldPath, newPath, err.Error())
				}
			}
		}
	}
}

type matchResult struct{
	base string
	index int
	ext string
}

func match (filename string) (*matchResult, error){
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])

	if err != nil{
		return nil, fmt.Errorf("%s did not match our pattern", filename)
	}
	return &matchResult{strings.Title(name), number, ext}, nil
}