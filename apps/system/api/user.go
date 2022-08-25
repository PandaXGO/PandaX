package api

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful/v3"
	"github.com/kakuilan/kgo"
	"github.com/mssola/user_agent"
	"log"

	"pandax/apps/system/api/form"
	"pandax/apps/system/api/vo"
	"pandax/apps/system/entity"

	logEntity "pandax/apps/log/entity"
	logServices "pandax/apps/log/services"

	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/captcha"
	filek "github.com/XM-GO/PandaKit/file"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/system/services"
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

// GenerateCaptcha 获取验证码
func (u *UserApi) GenerateCaptcha(request *restful.Request, response *restful.Response) {
	id, image := captcha.Generate()
	response.WriteEntity(vo.CaptchaVo{Base64Captcha: image, CaptchaId: id})
}

// RefreshToken 刷新token
func (u *UserApi) RefreshToken(rc *restfulx.ReqCtx) {
	tokenStr := rc.Request.Request.Header.Get("X-TOKEN")
	j := token.NewJWT("", []byte(global.Conf.Jwt.Key), jwt.SigningMethodHS256)
	token, err := j.RefreshToken(tokenStr)
	biz.ErrIsNil(err, "刷新token失败")
	rc.ResData = vo.TokenVo{
		Token:  token,
		Expire: time.Now().Unix() + global.Conf.Jwt.ExpireTime,
	}
}

// Login 用户登录
func (u *UserApi) Login(rc *restfulx.ReqCtx) {
	var l form.Login
	restfulx.BindQuery(rc, &l)
	log.Println(l)
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

	rc.ResData = vo.TokenVo{
		Token:  token,
		Expire: time.Now().Unix() + global.Conf.Jwt.ExpireTime,
	}
	go func() {
		var loginLog logEntity.LogLogin
		ua := user_agent.New(rc.Request.Request.UserAgent())
		loginLog.Ipaddr = rc.Request.Request.RemoteAddr
		loginLog.LoginLocation = utils.GetRealAddressByIP(rc.Request.Request.RemoteAddr)
		loginLog.LoginTime = time.Now()
		loginLog.Status = "0"
		loginLog.Remark = rc.Request.Request.UserAgent()
		browserName, browserVersion := ua.Browser()
		loginLog.Browser = browserName + " " + browserVersion
		loginLog.Os = ua.OS()
		loginLog.Platform = ua.Platform()
		loginLog.Username = login.Username
		loginLog.Msg = "登录成功"
		loginLog.CreateBy = login.Username
		u.LogLogin.Insert(loginLog)
	}()
}

// Auth 用户权限信息
func (u *UserApi) Auth(rc *restfulx.ReqCtx) {
	userName := restfulx.QueryParam(rc, "username")
	biz.NotEmpty(userName, "用户名必传")
	var user entity.SysUser
	user.Username = userName
	userData := u.UserApp.FindOne(user)
	role := u.RoleApp.FindOne(userData.RoleId)
	//前端权限
	permis := u.RoleMenuApp.GetPermis(role.RoleId)
	menus := u.MenuApp.SelectMenuRole(role.RoleKey)

	rc.ResData = vo.AuthVo{
		User:        *userData,
		Role:        *role,
		Permissions: permis,
		Menus:       Build(*menus),
	}
}

// LogOut 退出登录
func (u *UserApi) LogOut(rc *restfulx.ReqCtx) {
	var loginLog logEntity.LogLogin
	ua := user_agent.New(rc.Request.Request.UserAgent())
	loginLog.Ipaddr = rc.Request.Request.RemoteAddr
	loginLog.LoginTime = time.Now()
	loginLog.Status = "0"
	loginLog.Remark = rc.Request.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	loginLog.Browser = browserName + " " + browserVersion
	loginLog.Os = ua.OS()
	loginLog.Platform = ua.Platform()
	loginLog.Username = rc.LoginAccount.UserName
	loginLog.Msg = "退出成功"
	u.LogLogin.Insert(loginLog)
}

// GetSysUserList 列表数据
func (u *UserApi) GetSysUserList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	status := restfulx.QueryParam(rc, "status")
	username := restfulx.QueryParam(rc, "username")
	phone := restfulx.QueryParam(rc, "phone")

	deptId := restfulx.QueryInt(rc, "deptId", 0)
	var user entity.SysUser
	user.Status = status
	user.Username = username
	user.Phone = phone
	user.DeptId = int64(deptId)

	if !IsTenantAdmin(rc.LoginAccount.TenantId) {
		user.TenantId = rc.LoginAccount.TenantId
	}
	list, total := u.UserApp.FindListPage(pageNum, pageSize, user)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetSysUserProfile 获取当前登录用户
func (u *UserApi) GetSysUserProfile(rc *restfulx.ReqCtx) {

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

	rc.ResData = vo.UserProfileVo{
		Data:    user,
		PostIds: postIds,
		RoleIds: roleIds,
		Roles:   *roleList,
		Posts:   *postList,
		Dept:    *deptList,
	}
}

// InsetSysUserAvatar 修改头像
func (u *UserApi) InsetSysUserAvatar(rc *restfulx.ReqCtx) {
	form := rc.Request.Request.MultipartForm

	files := form.File["upload[]"]
	guid, _ := kgo.KStr.UuidV4()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		global.Log.Info(file.Filename)
		// 上传文件至指定目录
		biz.ErrIsNil(filek.SaveUploadedFile(file, filPath), "保存头像失败")
	}
	sysuser := entity.SysUser{}
	sysuser.UserId = rc.LoginAccount.UserId
	sysuser.Avatar = "/" + filPath
	sysuser.UpdateBy = rc.LoginAccount.UserName

	u.UserApp.Update(sysuser)
}

// SysUserUpdatePwd 修改密码
func (u *UserApi) SysUserUpdatePwd(rc *restfulx.ReqCtx) {
	var pws entity.SysUserPwd
	restfulx.BindQuery(rc, &pws)

	user := entity.SysUser{}
	user.UserId = rc.LoginAccount.UserId
	u.UserApp.SetPwd(user, pws)
}

// GetSysUser 获取用户
func (u *UserApi) GetSysUser(rc *restfulx.ReqCtx) {
	userId := restfulx.PathParamInt(rc, "userId")

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

	rc.ResData = vo.UserVo{
		Data:    result,
		PostIds: result.PostIds,
		RoleIds: result.RoleIds,
		Roles:   *roles,
		Posts:   *posts,
	}
}

// GetSysUserInit 获取添加用户角色和职位
func (u *UserApi) GetSysUserInit(rc *restfulx.ReqCtx) {

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
	rc.ResData = vo.UserRolePost{
		Roles: *roles,
		Posts: *posts,
	}
}

// GetUserRolePost 获取添加用户角色和职位
func (u *UserApi) GetUserRolePost(rc *restfulx.ReqCtx) {
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
	rc.ResData = vo.UserRolePost{
		Roles: roles,
		Posts: posts,
	}
}

// InsertSysUser 创建用户
func (u *UserApi) InsertSysUser(rc *restfulx.ReqCtx) {
	var sysUser entity.SysUser
	restfulx.BindQuery(rc, &sysUser)
	sysUser.CreateBy = rc.LoginAccount.UserName
	u.UserApp.Insert(sysUser)
}

// UpdateSysUser 修改用户数据
func (u *UserApi) UpdateSysUser(rc *restfulx.ReqCtx) {
	var sysUser entity.SysUser
	restfulx.BindQuery(rc, &sysUser)
	sysUser.CreateBy = rc.LoginAccount.UserName
	u.UserApp.Update(sysUser)
}

// UpdateSysUserStu 修改用户状态
func (u *UserApi) UpdateSysUserStu(rc *restfulx.ReqCtx) {
	var sysUser entity.SysUser
	restfulx.BindQuery(rc, &sysUser)
	sysUser.CreateBy = rc.LoginAccount.UserName
	u.UserApp.Update(sysUser)
}

// DeleteSysUser 删除用户数据
func (u *UserApi) DeleteSysUser(rc *restfulx.ReqCtx) {
	userIds := restfulx.PathParam(rc, "userId")
	u.UserApp.Delete(utils.IdsStrToIdsIntGroup(userIds))
}

// ExportUser 导出用户
func (u *UserApi) ExportUser(rc *restfulx.ReqCtx) {
	filename := restfulx.QueryParam(rc, "filename")
	status := restfulx.QueryParam(rc, "status")
	username := restfulx.QueryParam(rc, "username")
	phone := restfulx.QueryParam(rc, "phone")

	var user entity.SysUser
	user.Status = status
	user.Username = username
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
