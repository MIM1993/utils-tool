package time_utils

import (
	"time"
)

/**字符串->时间对象*/
func Str2Time(formatTimeStr string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime

}

/**字符串->时间戳*/
func Str2Stamp(formatTimeStr string) int64 {
	timeStruct := Str2Time(formatTimeStr)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

/**时间对象->字符串*/
const shortForm = "2006-01-01 15:04:05"
func Time2Str() string {

	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}

/*时间对象->时间戳*/
func Time2Stamp() int64 {
	t := time.Now()
	millisecond := t.UnixNano() / 1e6
	return millisecond
}

/*时间戳->字符串*/
func Stamp2Str(stamp int64) string {
	timeLayout := "2006-01-02 15:04:05"
	str := time.Unix(stamp/1000, 0).Format(timeLayout)
	return str
}

/*时间戳->时间对象*/
func Stamp2Time(stamp int64) time.Time {
	stampStr := Stamp2Str(stamp)
	timer := Str2Time(stampStr)
	return timer
}

//获取网络上的时间
func GetRemoteTime() (*time.Time, error) {
	var host string = "time.windows.com:123"
	// 182.92.12.11:123 是阿里的ntp服务器，可以换成其他域名的
	conn, err := net.Dial("udp", host)
	if err != nil {
		//log.Fatalf("failed to connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(15 * time.Second)); err != nil {
		//log.Fatalf("failed to set deadline: %v", err)
		return nil, err
	}
	req := &packet{Settings: 0x1B}
	if err := binary.Write(conn, binary.BigEndian, req); err != nil {
		//log.Fatalf("failed to send request: %v", err)
		return nil, err
	}
	rsp := &packet{}
	if err := binary.Read(conn, binary.BigEndian, rsp); err != nil {
		//log.Fatalf("failed to read server response: %v", err)
		return nil, err
	}
	secs := float64(rsp.TxTimeSec) - ntpEpochOffset
	nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32
	showtime := time.Unix(int64(secs), nanos)
	return &showtime, nil
}

