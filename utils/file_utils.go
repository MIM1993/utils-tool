package utils

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
)

//判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

//判断文件是否存在,不存在创建
func CreatFileIfNotExist(fileName string) error {
	_, err := os.Lstat(fileName)
	if !os.IsNotExist(err) {
		return nil
	}

	_, err = os.Create(fileName)
	return err
}

// CreateDirIfNotExists 如果文件夹目录不存在，创建文件夹及父目录
func CreateDirIfNotExists(path string, perm os.FileMode) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.MkdirAll(path, perm)
	}

	return err
}


// NewUUID 生成一个不含中划线的uuid
func NewUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

// GetSubAction 获取path的最后的一个路径
func GetSubAction(url string) (string, bool) {
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '.' {
			subAction := url[i+1:]
			return subAction, true
		}
	}
	return "", false
}

// GetSubPath 获取子路径
func GetSubPath(url, prefix string) (string, bool) {
	if strings.HasPrefix(url, prefix) {
		subPath := strings.TrimPrefix(url, prefix)
		return subPath, true
	}
	return "", false
}

// ProtectString 对敏感字符串进行保护
func ProtectString(str string) string {
	l := len(str)
	if l < 1 {
		return str
	}
	if l > 6 {
		l = 6
	}
	return str[0:l] + "***"
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
