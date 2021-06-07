package fileDelete

import (
	"io/ioutil"
	"fmt"
	"strings"
	"os"
)

func fileRead(filePath string)([]string){
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	lines := strings.Split(string(bytes),"\n")
	return lines
}

func deleteFile(lines []string){
	for line:= range lines{
		err := os.Remove(lines[line])
		if err != nil{
			fmt.Println(err)
		}
	}
}

func FileDelete(filePath string){
	defer deleteFile(fileRead(filePath))
}
