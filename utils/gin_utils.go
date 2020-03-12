package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

//gin中间件函数,用于设置一些基本对象,同时可以计算整体数据计算流程花费的时间
func logMiddleware(logger *logrus.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		reqID := c.Request.Header.Get("X-Request-Id")
		if reqID == "" {
			reqID = c.Request.Header.Get("request_id")
		}
		if reqID == "" {
			reqID = "随机字符串"
		}
		realIP := c.Request.Header.Get("X-Real-IP")

		logEntry := logger.WithFields(logrus.Fields{
			"request_id": reqID,
			"client":     c.Request.RemoteAddr,
			"real_ip":    realIP,
			"url_path":   c.Request.URL.Path,
		})

		c.Set("logEntry", logEntry)
		c.Set("request_id", reqID)
		var err error
		c.Set("error", err)

		begin := time.Now()
		c.Next()
		if cerr, ok := c.Get("error"); ok && cerr != nil {
			if ierr, ok := cerr.(error); ok || ierr == nil {
				err = ierr
			} else {
				err = fmt.Errorf("err: %v", cerr)
			}
		}

		logEntry.WithField("total_cost",
			time.Now().Sub(begin).Seconds()).Infof("finish. err:%v", err)
	}
}


//依据时间戳产生logid，该logid方便后续根据真实时间戳分类存储图文日志
func GenLogId() string {
	return fmt.Sprintf("%d.%d", time.Now().Unix(), rand.Intn(10000))
}

func ShortenStr(str string, length int) string {
	if len(str) <= length {
		return str
	}
	return str[:length] + "******"
}

