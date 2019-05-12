package ginhelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * format api response
 */
func JSONWrapper(f func(c *gin.Context) (int, string, gin.H)) func(*gin.Context) {
	return func(c *gin.Context) {
		code, msg, data := f(c)

		if code == 0 {
			msg = "success"
		}

		var resp = struct {
			Code int         `json:"code"`
			Msg  string      `json:"msg"`
			Data interface{} `json:"data"`
		}{
			Code: code,
			Msg:  msg,
			Data: data,
		}

		c.JSON(http.StatusOK, resp)
	}
}
