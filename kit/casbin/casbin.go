package casbin

import (
	"pandax/kit/biz"
	"pandax/kit/starter"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinService struct {
	ModelPath string
}

func (c *CasbinService) UpdateCasbin(roleKey string, casbinInfos []CasbinRule) error {
	c.ClearCasbin(0, roleKey)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{roleKey, v.Path, v.Method})
	}
	e := c.GetCasbinEnforcer()
	success, err := e.AddPolicies(rules)
	if err != nil {
		return err
	}
	if !success {
		return biz.NewBizErr("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func (c *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := starter.Db.Table("casbin_rule").Model(&CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	if err != nil {
		return biz.NewBizErr("修改api失败")
	}
	return nil
}

func (c *CasbinService) GetPolicyPathByRoleId(roleKey string) []CasbinRule {
	e := c.GetCasbinEnforcer()
	list := e.GetFilteredPolicy(0, roleKey)
	pathMaps := make([]CasbinRule, len(list))
	for i, v := range list {
		pathMaps[i] = CasbinRule{
			Path:   v[1],
			Method: v[2],
		}
	}
	return pathMaps
}

func (c *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := c.GetCasbinEnforcer()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (c *CasbinService) GetCasbinEnforcer() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(starter.Db)
		if err != nil {
			panic(biz.NewBizErr("新建权限适配器失败"))
		}
		syncedEnforcer, err = casbin.NewSyncedEnforcer(c.ModelPath, a)
		if err != nil {
			panic(biz.NewBizErr("新建权限适配器失败"))
		}
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
