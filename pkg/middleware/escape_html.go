package middleware

import (
	"github.com/emicklei/go-restful/v3"
	"html"
)

// 防止XSS攻击
func EscapeHTML(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// 获取请求参数中的HTML标签
	for _, p := range req.Request.URL.Query() {
		escaped := html.EscapeString(p[0])
		// 将转义后的参数重新设置到请求参数中
		req.Request.URL.Query().Set(p[0], escaped)
	}
	chain.ProcessFilter(req, resp)
}
