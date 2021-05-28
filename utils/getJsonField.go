import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

/**
 * 获取json字符串中指定字段内容  ioutil.ReadFile()读取字节切片
 * @param    bytes    json字符串字节数组
 * @param    field    可变参数，指定字段
 */
func getJsonField(bytes []byte, field ...string) []byte {
    if len(field) < 1 {
        fmt.Printf("At least two parameters are required.")
        return nil
    }

    //将字节切片映射到指定map上  key：string类型，value：interface{}  类型能存任何数据类型
    var mapObj map[string]interface{}
    json.Unmarshal(bytes, &mapObj)
    var tmpObj interface{}
    tmpObj = mapObj
    for i := 0; i < len(field); i++ {
        tmpObj = tmpObj.(map[string]interface{})[field[i]]
        if tmpObj == nil {
            fmt.Printf("No field specified: %s ", field[i])
            return nil
        }
    }

    result, err := json.Marshal(tmpObj)
    if err != nil {
        fmt.Print(err)
        return nil
    }
    return result
}

func main() {
    bytes, _ := ioutil.ReadFile("./data.json")
    s := getJsonField(bytes, "basic_info", "tss")
    fmt.Println(string(s))
}
