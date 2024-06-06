package api

import (
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/kit/biz"
	"pandax/kit/restfulx"
	"pandax/kit/utils"
	"strings"
)

type ProductCategoryApi struct {
	ProductCategoryApp services.ProductCategoryModel
}

// GetProductCategoryTree ProductCategory 树
func (p *ProductCategoryApi) GetProductCategoryTree(rc *restfulx.ReqCtx) {
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := restfulx.QueryParam(rc, "id")
	/*err := rc.Validate.Var(id, "required")
	biz.ErrIsNil(err, "id 必传")*/
	sg := entity.ProductCategory{Name: name, Status: status}
	sg.Id = id
	rc.ResData = p.ProductCategoryApp.SelectProductCategory(sg)
}

func (p *ProductCategoryApi) GetProductCategoryTreeLabel(rc *restfulx.ReqCtx) {
	sg := entity.ProductCategory{}
	rc.ResData = p.ProductCategoryApp.SelectProductCategoryLabel(sg)
}

func (p *ProductCategoryApi) GetProductCategoryList(rc *restfulx.ReqCtx) {
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := restfulx.QueryParam(rc, "id")

	sg := entity.ProductCategory{Name: name, Status: status}
	sg.Id = id
	if sg.Name == "" {
		rc.ResData = p.ProductCategoryApp.SelectProductCategory(sg)
	} else {
		list, err := p.ProductCategoryApp.FindList(sg)
		biz.ErrIsNil(err, "查询产品分类列表失败")
		rc.ResData = list
	}
}

// GetProductCategoryAllList 查询所有
func (p *ProductCategoryApi) GetProductCategoryAllList(rc *restfulx.ReqCtx) {
	var vsg entity.ProductCategory
	list, err := p.ProductCategoryApp.FindList(vsg)
	biz.ErrIsNil(err, "查询产品分类列表失败")
	rc.ResData = list
}

// GetProductCategory 获取ProductCategory
func (p *ProductCategoryApi) GetProductCategory(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one, err := p.ProductCategoryApp.FindOne(id)
	biz.ErrIsNil(err, "查询产品分类失败")
	rc.ResData = one
}

// InsertProductCategory 添加ProductCategory
func (p *ProductCategoryApi) InsertProductCategory(rc *restfulx.ReqCtx) {
	var data entity.ProductCategory
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = utils.GenerateID()
	data.Owner = rc.LoginAccount.UserName
	data.OrgId = rc.LoginAccount.OrganizationId
	p.ProductCategoryApp.Insert(data)
}

// UpdateProductCategory 修改ProductCategory
func (p *ProductCategoryApi) UpdateProductCategory(rc *restfulx.ReqCtx) {
	var data entity.ProductCategory
	restfulx.BindJsonAndValid(rc, &data)

	p.ProductCategoryApp.Update(data)
}

// DeleteProductCategory 删除ProductCategory
func (p *ProductCategoryApi) DeleteProductCategory(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.ProductCategoryApp.Delete(ids)
}
