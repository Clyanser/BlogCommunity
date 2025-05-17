package middleware

import (
	"GoBlog/service/log_service"
	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.Body = append(rw.Body, b...)
	return rw.ResponseWriter.Write(b)
}

func LogMiddleware(c *gin.Context) {
	log := log_service.NewActionLog(c)
	log.SetRequestBody(c)

	res := &ResponseWriter{
		ResponseWriter: c.Writer,
	}
	c.Writer = res
	c.Next()
	//响应中间件
	log.SetResponseBody(res.Body)
	//log.Save()
}
