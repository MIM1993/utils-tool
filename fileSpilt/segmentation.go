package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

//大文件分割
const chunkSize = 4 << 20



func main() {
	SplitFile("fileSpilt/test.mp4")

	merge("fileSpilt/test", "new.mp4")
}

//分割
func SplitFile(fileName string) {
	fmt.Println("begin ...")
	fn := fmt.Sprintf("%s", fileName)
	fileInfo, err := os.Stat(fn)
	if err != nil {
		fmt.Println("stat err:", err)
		return
	}

	//计算分包数量
	tempNum := float64(fileInfo.Size()) / float64(chunkSize)
	num := int(math.Ceil(tempNum))

	//读取分割文件
	f, err := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	defer f.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

	dir := fmt.Sprintf("%s", strings.Split(fileName, ".")[0])
	fmt.Println(dir)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("mkdir err:", err)
		return
	}

	b := make([]byte, chunkSize)
	var i int64 = 1
	for ; i <= int64(num); i++ {
		//设置偏移值
		f.Seek((i-1)*chunkSize, 0)
		if len(b) > int(fileInfo.Size()-(i-1)*chunkSize) {
			b = make([]byte, fileInfo.Size()-(i-1)*chunkSize)
		}

		_, err := f.Read(b)
		if err != nil {
			fmt.Println("read chunk err:", err)
			return
		}

		chunkName := fmt.Sprintf("%s/%s.db", dir, strconv.FormatInt(i, 10))
		fck, err := os.OpenFile(chunkName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			fmt.Println("open fck err:", err)
			return
		}
		fck.Write(b)
		fck.Close()
	}

	fmt.Println("finished ...")

}

//合并
func merge(dirName, newFileName string) {
	fmt.Println("begin ...")
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("read dir err:", err)
		return
	}
	fn := fmt.Sprintf("fileSpilt/%s", newFileName)
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer f.Close()
	if err != nil {
		fmt.Println("openFile err:", err)
		return
	}

	for i := 1; i <= len(files); i++ {
		tmpName := fmt.Sprintf("%s/%d.db", dirName, i)
		fmt.Println(tmpName)
		tf, err := os.OpenFile(tmpName, os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println("openFile db err:", err)
			return
		}
		data, err := ioutil.ReadAll(tf)
		if err != nil {
			fmt.Println("ReadAll db err:", err)
			return
		}
		f.Write(data)
		tf.Close()
	}

	os.RemoveAll(dirName)

	fmt.Println("finished ...")
}
