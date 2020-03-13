package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestFileExist(t *testing.T){
	if FileExist("hello"){
		fmt.Println("file exist")
	}else {
		fmt.Println("file Not exist")
	}
}

func TestCreatFileIfNotExist(t *testing.T){
	err := CreatFileIfNotExist("./hello/hello.txt")

	if err!=nil{
		fmt.Println(err)
	}
}

func TestCreateDirIfNotExists(t *testing.T) {
	err :=CreateDirIfNotExists("./hello",os.ModePerm)
	if err!=nil{
		fmt.Println(err)
	}
}