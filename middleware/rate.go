package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"pandax/base/config"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/19 8:28
 **/

//限流中间件
func Rate() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(config.Conf.Server.Rate.RateNum, nil)
	lmt.SetMessage("已经超出接口请求限制，稍后再试.")
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
