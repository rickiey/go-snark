package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// response 给前端响应的结构
type response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Result response result
func Result(c *gin.Context, httpCode int, ecode int, msg string, data interface{}) {
	c.Set(ContextErrCode, ecode)
	c.JSON(httpCode, response{
		Code:    ecode,
		Message: msg,
		Data:    data,
	})
}

// InnerErr 500
// from 错误日志来源
func InnerErr(c *gin.Context, from string, err error) {
	glog.Infof("%s: %s", from, err.Error())
	Result(c, http.StatusInternalServerError, Failed, FailedMsg, nil)
}

// NormalInnerErr 内部请求错误，转换给前台正常业务error, 并记录日志
func NormalInnerErr(c *gin.Context, ecode int, msg string, from string, err error) {
	glog.Infof("%s: %s", from, err.Error())
	Result(c, http.StatusOK, ecode, msg, nil)
}

// NormalErr 业务error
func NormalErr(c *gin.Context, ecode int, msg string) {
	Result(c, http.StatusOK, ecode, msg, nil)
}

// NormalResult 正常结果
func NormalResult(c *gin.Context, data interface{}) {
	Result(c, http.StatusOK, OK, OkMsg, data)
}

// OutputJSON 直接输出json字符串
func OutputJSON(c *gin.Context, jsonStr string) {
	c.Data(http.StatusOK, gin.MIMEJSON, []byte(`{"code": 0, "msg": "SUCCESS", "data": `+jsonStr+`}`))
}
