package middleware

import (
	"go-snark/resp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// LoggerHandler ..
func LoggerHandler(c *gin.Context) {
	// Start timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	method := c.Request.Method

	// Process request
	c.Next()

	// Stop timer
	end := time.Now()
	latency := end.Sub(start)
	statusCode := c.Writer.Status()
	ecode := c.GetInt(resp.ContextErrCode)
	clientIP := c.ClientIP()
	if "" != raw {
		path = path + "?" + raw
	}
	if path == "/ping" || method == "GET" {
		return
	}
	glog.Infof("METHOD:%s | PATH:%s | CODE:%d | IP:%s | TIME:%d | ECODE:%d", method, path, statusCode, clientIP,
		latency/time.Millisecond, ecode)
}
