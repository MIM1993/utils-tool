package internal

import (
	"reflect"
	"unsafe"
	"github.com/axgle/mahonia"
)

// BytesToString converts byte slice to a string without memory allocation.
//
// Note it may break if the implementation of string or slice header changes in the future go versions.
func BytesToString(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts string to a byte slice without memory allocation.
//
// Note it may break if the implementation of string or slice header changes in the future go versions.
func StringToBytes(s string) (b []byte) {
	/* #nosec G103 */
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

//针对 json 解析的特性,对reflect.DeepEqual 函数进行改造适配。
func DeepEqual(v1, v2 interface{}) bool {
	 if reflect.DeepEqual(v1, v2) {
	  return true
	 }
	 bytesA, _ := json.Marshal(v1)
	 bytesB, _ := json.Marshal(v2)
	 return bytes.Equal(bytesA, bytesB)
}

//将[]byte形式的数字转为int
func byteToInt(number []byte) (int, error) {
	var sum int
	x := 1
	for i := len(number) - 1; i >= 0; i-- {
		if number[i] < '0' || number[i] > '9' {
			return 0, errors.New("toNumberError")
		}
		sum += int(number[i]-'0') * x
		x = x * 10
	}

	return sum, nil
}

//将gbk格式转换为utf8格式
func CoverGBKToUTF8(src string) string {
	return mahonia.NewDecoder("gbk").ConvertString(src)
}
