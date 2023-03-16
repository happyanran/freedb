package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/happyanran/freedb/server/global"

	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

type CasbinService struct{}

func (casbinService *CasbinService) Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.FDB_DB)
		if err != nil {
			global.FDB_LOG.Errorf("适配数据库失败请检查casbin表是否为InnoDB引擎! %v", err)
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			global.FDB_LOG.Errorf("字符串加载模型失败! %v", err)
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(600)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}
