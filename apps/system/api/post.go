package api

import (
	"errors"
	"fmt"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"pandax/kit/biz"
	"pandax/kit/model"
	"pandax/kit/restfulx"
	"pandax/kit/utils"
	"pandax/pkg/global"
)

type PostApi struct {
	PostApp services.SysPostModel
	UserApp services.SysUserModel
	RoleApp services.SysRoleModel
}

// GetPostList 职位列表数据
func (p *PostApi) GetPostList(rc *restfulx.ReqCtx) {

	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	status := restfulx.QueryParam(rc, "status")
	postName := restfulx.QueryParam(rc, "postName")
	postCode := restfulx.QueryParam(rc, "postCode")
	post := entity.SysPost{Status: status, PostName: postName, PostCode: postCode}

	list, total, err := p.PostApp.FindListPage(pageNum, pageSize, post)
	biz.ErrIsNil(err, "查询部门列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetPost 获取职位
func (p *PostApi) GetPost(rc *restfulx.ReqCtx) {
	postId := restfulx.PathParamInt(rc, "postId")
	data, err := p.PostApp.FindOne(int64(postId))
	biz.ErrIsNil(err, "查询部门失败")
	rc.ResData = data
}

// InsertPost 添加职位
func (p *PostApi) InsertPost(rc *restfulx.ReqCtx) {
	var post entity.SysPost
	restfulx.BindJsonAndValid(rc, &post)
	post.CreateBy = rc.LoginAccount.UserName
	_, err := p.PostApp.Insert(post)
	biz.ErrIsNil(err, "添加部门失败")
}

// UpdatePost 修改职位
func (p *PostApi) UpdatePost(rc *restfulx.ReqCtx) {
	var post entity.SysPost
	restfulx.BindJsonAndValid(rc, &post)

	post.CreateBy = rc.LoginAccount.UserName
	err := p.PostApp.Update(post)
	biz.ErrIsNil(err, "修改部门失败")
}

// DeletePost 删除职位
func (p *PostApi) DeletePost(rc *restfulx.ReqCtx) {
	postId := restfulx.PathParam(rc, "postId")
	postIds := utils.IdsStrToIdsIntGroup(postId)

	deList := make([]int64, 0)
	for _, id := range postIds {
		user := entity.SysUser{}
		user.PostId = id
		list, err := p.UserApp.FindList(user)
		if err != nil {
			continue
		}
		if len(*list) == 0 {
			deList = append(deList, id)
		} else {
			global.Log.Info(fmt.Sprintf("dictId: %d 存在岗位绑定用户无法删除", id))
		}
	}
	if len(deList) == 0 {
		biz.ErrIsNil(errors.New("所有岗位都已绑定用户，无法删除"), "所有岗位都已绑定用户，无法删除")
	}
	err := p.PostApp.Delete(deList)
	biz.ErrIsNil(err, "删除部门失败")
}
