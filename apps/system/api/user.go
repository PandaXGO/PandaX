package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kakuilan/kgo"
	"github.com/mssola/user_agent"
	"net/http"
	"pandax/apps/system/api/form"
	"pandax/apps/system/api/vo"
	"pandax/apps/system/entity"
	"pandax/base/token"

	logEntity "pandax/apps/log/entity"
	logServices "pandax/apps/log/services"

	"pandax/apps/system/services"
	"pandax/base/biz"
	"pandax/base/captcha"
	"pandax/base/ginx"
	"pandax/base/utils"
	"pandax/pkg/global"
	"strings"
	"time"
)

type UserApi struct {
	UserApp     services.SysUserModel
	MenuApp     services.SysMenuModel
	PostApp     services.SysPostModel
	RoleApp     services.SysRoleModel
	RoleMenuApp services.SysRoleMenuModel
	DeptApp     services.SysDeptModel
	LogLogin    logServices.LogLoginModel
}

// @Tags Base
// @Summary 获取验证码
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/user/getCaptcha [get]
func (u *UserApi) GenerateCaptcha(c *gin.Context) {
	id, image := captcha.Generate()
	c.JSON(http.StatusOK, map[string]any{"base64Captcha": image, "captchaId": id})
}

// @Tags Base
// @Summary 刷新token
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/user/refreshToken [get]
func (u *UserApi) RefreshToken(rc *ginx.ReqCtx) {
	tokenStr := rc.GinCtx.Request.Header.Get("X-TOKEN")
	j := token.NewJWT("", []byte(global.Conf.Jwt.Key), jwt.SigningMethodHS256)
	token, err := j.RefreshToken(tokenStr)
	biz.ErrIsNil(err, "刷新token失败")
	rc.ResData = map[string]any{"token": token, "expire": time.Now().Unix() + global.Conf.Jwt.ExpireTime}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body form.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/user/login [post]
func (u *UserApi) Login(rc *ginx.ReqCtx) {
	var l form.Login
	ginx.BindJsonAndValid(rc.GinCtx, &l)
	biz.IsTrue(captcha.Verify(l.CaptchaId, l.Captcha), "验证码认证失败")

	login := u.UserApp.Login(entity.Login{Username: l.Username, Password: l.Password})
	role := u.RoleApp.FindOne(login.RoleId)
	j := token.NewJWT("", []byte(global.Conf.Jwt.Key), jwt.SigningMethodHS256)
	token, err := j.CreateToken(token.Claims{
		UserId:   login.UserId,
		TenantId: login.TenantId,
		UserName: login.Username,
		RoleId:   login.RoleId,
		RoleKey:  role.RoleKey,
		DeptId:   login.DeptId,
		PostId:   login.PostId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                       // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.Conf.Jwt.ExpireTime, // 过期时间 7天  配置文件
			Issuer:    "PandaX",                                       // 签名的发行者
		},
	})
	biz.ErrIsNil(err, "生成Token失败")

	rc.ResData = map[string]any{
		"token":  token,
		"expire": time.Now().Unix() + global.Conf.Jwt.ExpireTime,
	}

	var loginLog logEntity.LogLogin
	ua := user_agent.New(rc.GinCtx.Request.UserAgent())
	loginLog.Ipaddr = rc.GinCtx.ClientIP()
	loginLog.LoginLocation = utils.GetRealAddressByIP(rc.GinCtx.ClientIP())
	loginLog.LoginTime = time.Now()
	loginLog.Status = "0"
	loginLog.Remark = rc.GinCtx.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	loginLog.Browser = browserName + " " + browserVersion
	loginLog.Os = ua.OS()
	loginLog.Platform = ua.Platform()
	loginLog.Username = login.Username
	loginLog.Msg = "登录成功"
	loginLog.CreateBy = login.Username
	u.LogLogin.Insert(loginLog)
}

// @Tags Base
// @Summary 用户权限信息
// @Param userName query string false "userName"
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /system/user/auth [get]
func (u *UserApi) Auth(rc *ginx.ReqCtx) {
	userName := rc.GinCtx.Query("username")
	biz.NotEmpty(userName, "用户名必传")
	var user entity.SysUser
	user.Username = userName
	userData := u.UserApp.FindOne(user)
	role := u.RoleApp.FindOne(userData.RoleId)
	//前端权限
	permis := u.RoleMenuApp.GetPermis(role.RoleId)
	menus := u.MenuApp.SelectMenuRole(role.RoleKey)

	rc.ResData = map[string]any{
		"user":        userData,
		"role":        role,
		"permissions": permis,
		"menus":       Build(*menus),
	}
}

// @Tags Base
// @Summary 退出登录
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /system/user/logout [post]
func (u *UserApi) LogOut(rc *ginx.ReqCtx) {
	var loginLog logEntity.LogLogin
	ua := user_agent.New(rc.GinCtx.Request.UserAgent())
	loginLog.Ipaddr = rc.GinCtx.ClientIP()
	loginLog.LoginTime = time.Now()
	loginLog.Status = "0"
	loginLog.Remark = rc.GinCtx.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	loginLog.Browser = browserName + " " + browserVersion
	loginLog.Os = ua.OS()
	loginLog.Platform = ua.Platform()
	loginLog.Username = rc.LoginAccount.UserName
	loginLog.Msg = "退出成功"
	u.LogLogin.Insert(loginLog)
}

// @Summary 列表数据
// @Description 获取JSON
// @Tags 用户
// @Param userName query string false "userName"
// @Param phone query string false "phone"
// @Param status query string false "status"
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /system/user/sysUserList [get]
// @Security X-TOKEN
func (u *UserApi) GetSysUserList(rc *ginx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	status := rc.GinCtx.Query("status")
	userName := rc.GinCtx.Query("username")
	phone := rc.GinCtx.Query("phone")
	deptId := ginx.QueryInt(rc.GinCtx, "deptId", 0)
	var user entity.SysUser
	user.Status = status
	user.Username = userName
	user.Phone = phone
	user.DeptId = int64(deptId)

	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		user.TenantId = rc.LoginAccount.TenantId
	}
	list, total := u.UserApp.FindListPage(pageNum, pageSize, user)

	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 获取当前登录用户
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/user/profile [get]
// @Security
func (u *UserApi) GetSysUserProfile(rc *ginx.ReqCtx) {

	sysUser := entity.SysUser{}
	sysUser.UserId = rc.LoginAccount.UserId
	user := u.UserApp.FindOne(sysUser)

	//获取角色列表
	roleList := u.RoleApp.FindList(entity.SysRole{RoleId: rc.LoginAccount.RoleId})
	//岗位列表
	postList := u.PostApp.FindList(entity.SysPost{PostId: rc.LoginAccount.PostId})
	//获取部门列表
	deptList := u.DeptApp.FindList(entity.SysDept{DeptId: rc.LoginAccount.DeptId})

	postIds := make([]int64, 0)
	postIds = append(postIds, rc.LoginAccount.PostId)

	roleIds := make([]int64, 0)
	roleIds = append(roleIds, rc.LoginAccount.RoleId)

	rc.ResData = map[string]any{
		"data":    user,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roleList,
		"posts":   postList,
		"dept":    deptList,
	}
}

// @Summary 修改头像
// @Description 修改头像
// @Tags 用户
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /system/user/profileAvatar [post]
func (u *UserApi) InsetSysUserAvatar(rc *ginx.ReqCtx) {
	form, err := rc.GinCtx.MultipartForm()
	biz.ErrIsNil(err, "头像上传失败")

	files := form.File["upload[]"]
	guid, _ := kgo.KStr.UuidV4()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		global.Log.Info(file.Filename)
		// 上传文件至指定目录
		biz.ErrIsNil(rc.GinCtx.SaveUploadedFile(file, filPath), "保存头像失败")
	}
	sysuser := entity.SysUser{}
	sysuser.UserId = rc.LoginAccount.UserId
	sysuser.Avatar = "/" + filPath
	sysuser.UpdateBy = rc.LoginAccount.UserName

	u.UserApp.Update(sysuser)
}

// @Summary 修改密码
// @Description 修改密码
// @Tags 用户
// @Param pwd body entity.SysUserPwd true "pwd"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /system/user/updatePwd [post]
func (u *UserApi) SysUserUpdatePwd(rc *ginx.ReqCtx) {
	var pws entity.SysUserPwd
	ginx.BindJsonAndValid(rc.GinCtx, &pws)

	user := entity.SysUser{}
	user.UserId = rc.LoginAccount.UserId
	u.UserApp.SetPwd(user, pws)
}

// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/user/sysUser/{userId} [get]
// @Security
func (u *UserApi) GetSysUser(rc *ginx.ReqCtx) {
	userId := ginx.PathParamInt(rc.GinCtx, "userId")

	user := entity.SysUser{}
	user.UserId = int64(userId)
	result := u.UserApp.FindOne(user)

	var role entity.SysRole
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		role.TenantId = rc.LoginAccount.TenantId
	}
	roles := u.RoleApp.FindList(role)

	var post entity.SysPost
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		post.TenantId = rc.LoginAccount.TenantId
	}
	posts := u.PostApp.FindList(post)

	rc.ResData = map[string]any{
		"data":    result,
		"postIds": result.PostIds,
		"roleIds": result.RoleIds,
		"roles":   roles,
		"posts":   posts,
	}
}

// @Summary 获取添加用户角色和职位
// @Description 获取JSON
// @Tags 用户
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/user/getInit [get]
// @Security
func (u *UserApi) GetSysUserInit(rc *ginx.ReqCtx) {

	var role entity.SysRole
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		role.TenantId = rc.LoginAccount.TenantId
	}
	roles := u.RoleApp.FindList(role)
	var post entity.SysPost
	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		post.TenantId = rc.LoginAccount.TenantId
	}
	posts := u.PostApp.FindList(post)
	mp := make(map[string]any, 2)
	mp["roles"] = roles
	mp["posts"] = posts
	rc.ResData = mp
}

// @Summary 获取添加用户角色和职位
// @Description 获取JSON
// @Tags 用户
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/user/getInit [get]
// @Security
func (u *UserApi) GetUserRolePost(rc *ginx.ReqCtx) {
	var user entity.SysUser
	user.UserId = rc.LoginAccount.UserId

	resData := u.UserApp.FindOne(user)

	roles := make([]entity.SysRole, 0)
	posts := make([]entity.SysPost, 0)
	for _, roleId := range strings.Split(resData.RoleIds, ",") {
		ro := u.RoleApp.FindOne(kgo.KConv.Str2Int64(roleId))
		roles = append(roles, *ro)
	}
	for _, postId := range strings.Split(resData.PostIds, ",") {
		po := u.PostApp.FindOne(kgo.KConv.Str2Int64(postId))
		posts = append(posts, *po)
	}
	mp := make(map[string]any, 2)
	mp["roles"] = roles
	mp["posts"] = posts
	rc.ResData = mp
}

// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/user/sysUser [post]
func (u *UserApi) InsertSysUser(rc *ginx.ReqCtx) {
	var sysUser entity.SysUser
	ginx.BindJsonAndValid(rc.GinCtx, &sysUser)
	sysUser.CreateBy = rc.LoginAccount.UserName
	u.UserApp.Insert(sysUser)
}

// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysUser true "用户数据"g
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/user/sysUser [put]
func (u *UserApi) UpdateSysUser(rc *ginx.ReqCtx) {
	var sysUser entity.SysUser
	ginx.BindJsonAndValid(rc.GinCtx, &sysUser)
	sysUser.CreateBy = rc.LoginAccount.UserName
	u.UserApp.Update(sysUser)
}

// @Summary 修改用户状态
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/user/sysUser [put]
func (u *UserApi) UpdateSysUserStu(rc *ginx.ReqCtx) {
	var sysUser entity.SysUser
	ginx.BindJsonAndValid(rc.GinCtx, &sysUser)
	sysUser.CreateBy = rc.LoginAccount.UserName
	u.UserApp.Update(sysUser)
}

// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "多个id 使用逗号隔开"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/user/sysuser/{userId} [delete]
func (u *UserApi) DeleteSysUser(rc *ginx.ReqCtx) {
	userIds := rc.GinCtx.Param("userId")
	us := utils.IdsStrToIdsIntGroup(userIds)
	u.UserApp.Delete(us)
}

// @Summary 导出用户
// @Description 导出数据
// @Tags 用户
// @Param userName query string false "userName"
// @Param phone query string false "phone"
// @Param status query string false "status"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/dict/type/export [get]
func (u *UserApi) ExportUser(rc *ginx.ReqCtx) {
	filename := rc.GinCtx.Query("filename")
	status := rc.GinCtx.Query("status")
	userName := rc.GinCtx.Query("username")
	phone := rc.GinCtx.Query("phone")

	var user entity.SysUser
	user.Status = status
	user.Username = userName
	user.Phone = phone

	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		user.TenantId = rc.LoginAccount.TenantId
	}

	list := u.UserApp.FindList(user)
	fileName := utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, fileName)
	rc.Download(fileName)
}

// Build 构建前端路由
func Build(menus []entity.SysMenu) []vo.RouterVo {
	equals := func(a string, b string) bool {
		if a == b {
			return true
		}
		return false
	}
	rvs := make([]vo.RouterVo, 0)
	for _, ms := range menus {
		var rv vo.RouterVo
		rv.Name = ms.Path
		rv.Path = ms.Path
		rv.Component = ms.Component
		auth := make([]string, 0)
		if ms.Permission != "" {
			auth = strings.Split(ms.Permission, ",")
		}
		rv.Meta = vo.MetaVo{
			Title:       ms.MenuName,
			IsLink:      ms.IsLink,
			IsHide:      equals("1", ms.IsHide),
			IsKeepAlive: equals("0", ms.IsKeepAlive),
			IsAffix:     equals("0", ms.IsAffix),
			IsIframe:    equals("0", ms.IsIframe),
			Auth:        auth,
			Icon:        ms.Icon,
		}
		rv.Children = Build(ms.Children)
		rvs = append(rvs, rv)
	}

	return rvs
}
