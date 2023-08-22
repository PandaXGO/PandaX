package api

import (
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
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
		rc.ResData = p.ProductCategoryApp.FindList(sg)
	}
}

// GetProductCategoryAllList 查询所有
func (p *ProductCategoryApi) GetProductCategoryAllList(rc *restfulx.ReqCtx) {
	var vsg entity.ProductCategory
	rc.ResData = p.ProductCategoryApp.FindList(vsg)
}

// GetProductCategory 获取ProductCategory
func (p *ProductCategoryApi) GetProductCategory(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.ProductCategoryApp.FindOne(id)
}

// InsertProductCategory 添加ProductCategory
func (p *ProductCategoryApi) InsertProductCategory(rc *restfulx.ReqCtx) {
	var data entity.ProductCategory
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = kgo.KStr.Uniqid("pc_")
	data.Owner = rc.LoginAccount.UserName
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
