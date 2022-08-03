package middleware

import (
	"github.com/emicklei/go-restful/v3"
)

// 处理跨域请求,支持options访问
func Cors(wsContainer *restful.Container) restful.CrossOriginResourceSharing {
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowedHeaders: []string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token", "X-Token", "X-User-Id"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "DELETE", "PUT"},
		CookiesAllowed: false,
		Container:      wsContainer}

	return cors
}
