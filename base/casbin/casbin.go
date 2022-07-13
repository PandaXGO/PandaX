package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"pandax/base/biz"
	"pandax/base/config"
	"pandax/base/global"
	"sync"
)

func UpdateCasbin(roleKey string, casbinInfos []CasbinRule) error {
	ClearCasbin(0, roleKey)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{roleKey, v.Path, v.Method})
	}
	e := Casbin()
	success, _ := e.AddPolicies(rules)
	biz.IsTrue(success, "存在相同api,添加失败,请联系管理员")
	return nil
}

func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) {
	err := global.Db.Table("casbin_rule").Model(&CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]any{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	biz.ErrIsNil(err, "修改api失败")
}

func GetPolicyPathByRoleId(roleKey string) (pathMaps []CasbinRule) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, roleKey)
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinRule{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.Db)
		biz.ErrIsNil(err, "新建权限适配器失败")
		syncedEnforcer, err = casbin.NewSyncedEnforcer(config.Conf.Casbin.ModelPath, a)
		biz.ErrIsNil(err, "新建权限适配器失败")
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
