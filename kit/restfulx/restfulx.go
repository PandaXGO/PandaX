package restfulx

import (
	"encoding/json"
	"net/http"
	"pandax/kit/biz"
	"pandax/kit/logger"
	"pandax/kit/model"
	"strconv"

	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
)

// 绑定并校验请求结构体参数  结构体添加 例如： validate:"required" validate:"required,gt=10"
func BindJsonAndValid(rc *ReqCtx, data any) {
	if err := rc.Request.ReadEntity(data); err != nil {
		panic(any(biz.NewBizErr(err.Error())))
	}
	if err := rc.Validate.Struct(data); err != nil {
		panic(any(biz.NewBizErr("传参格式错误：" + err.Error())))
	}
}

// BindQuery 绑定查询字符串到
func BindQuery(rc *ReqCtx, data any) {
	if err := rc.Request.ReadEntity(data); err != nil {
		panic(any(biz.NewBizErr(err.Error())))
	}
}

// PathParamsToAny 获取虽有路径中的参数
func PathParamsToAny(rc *ReqCtx, in any) {
	vars := make(map[string]any)
	for k, v := range rc.Request.PathParameters() {
		vars[k] = v
	}
	marshal, _ := json.Marshal(vars)
	err := json.Unmarshal(marshal, in)
	biz.ErrIsNil(err, "error get path value encoding unmarshal")
	return
}

// QueryParamsToAny 获取所有Query的参数
func QueryParamsToAny(rc *ReqCtx, in any) {
	err := rc.Request.Request.ParseForm()
	biz.ErrIsNil(err, "error get ParseForm value encoding unmarshal")
	decoder := schema.NewDecoder()
	err = decoder.Decode(in, rc.Request.Request.Form)
	biz.ErrIsNil(err, "error get path value encoding unmarshal")
	return
}

// GetPageQueryParam 获取分页参数
func GetPageQueryParam(rc *ReqCtx) *model.PageParam {
	return &model.PageParam{PageNum: QueryInt(rc, "pageNum", 1), PageSize: QueryInt(rc, "pageSize", 10)}
}

// QueryInt 获取查询参数中指定参数值，并转为int
func QueryInt(rc *ReqCtx, qm string, defaultInt int) int {
	qv := rc.Request.QueryParameter(qm)
	if qv == "" {
		return defaultInt
	}
	qvi, err := strconv.Atoi(qv)
	biz.ErrIsNil(err, "query param not int")
	return qvi
}

// QueryParam QueryParam
func QueryParam(rc *ReqCtx, key string) string {
	return rc.Request.QueryParameter(key)
}

// PathParamInt 获取路径参数
func PathParamInt(rc *ReqCtx, pm string) int {
	value, _ := strconv.Atoi(rc.Request.PathParameter(pm))
	return value
}
func PathParam(rc *ReqCtx, pm string) string {
	return rc.Request.PathParameter(pm)
}

// 文件下载
func Download(rc *ReqCtx, filename string) {
	rc.Response.Header().Add("success", "true")
	rc.Response.Header().Set("Content-Length", "-1")
	rc.Response.Header().Set("Content-Disposition", "attachment; filename="+filename)
	http.ServeFile(
		rc.Response.ResponseWriter,
		rc.Request.Request, filename)
}

// 返回统一成功结果
func SuccessRes(response *restful.Response, data any) {
	response.WriteEntity(model.Success(data))
}

// 返回失败结果集
func ErrorRes(response *restful.Response, err any) {
	switch t := err.(type) {
	case *biz.BizError:
		response.WriteEntity(model.Error(t))
		break
	case error:
		response.WriteEntity(model.ServerError())
		logger.Log.Error(t)
		break
	case string:
		response.WriteEntity(model.ServerError())
		logger.Log.Error(t)
		break
	default:
		logger.Log.Error(t)
	}
}
