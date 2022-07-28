package system

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"miya/api/internal/model"
)

func (s SystemLogic) EnforcerTool(ctx context.Context) *casbin.Enforcer {
	//TODO implement me
	//adapter := gormadapter.NewAdapterByDB(s.UserActivityRepo.)
	db := s.CasBinRepo.GetDb(ctx)
	a, err := gormadapter.NewAdapterByDBWithCustomTable(db, &model.CasbinRule{})
	fmt.Println("NewAdapterByDBWithCustomTable", err)

	e, err := casbin.NewEnforcer("etc/casbin.conf", a)
	fmt.Println("NewEnforcer", err)
	err = e.LoadPolicy()
	fmt.Println("LoadPolicy", err)
	// Check the permission.
	//e.AddPolicy("admin", "domain", "data1", "read")
	// e.RemovePolicy("admin", "domain", "data1", "read")
	m, h := e.Enforce("admin", "domain", "data1", "read")
	fmt.Println("Enforce", m, h)
	save := e.SavePolicy()
	fmt.Println("SavePolicy", save)
	return e
}

//LoadPolicy()	强制的	从存储中加载所有策略规则
//SavePolicy()	强制的	将所有策略规则保存到存储中
//AddPolicy()	可选择的	向存储中添加策略规则
//RemovePolicy()	可选择的	从存储中删除策略规则
//RemoveFilteredPolicy()	可选择的	从存储中删除匹配筛选器的策略规则

// AddPolicy 向存储器添加了一条策略规则。
//func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
//	return errors.New("not implemented")
//}
//
//// RemovePolicy 从存储器中移除一条策略规则。
//func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
//	return errors.New("not implemented")
//}
//
//// RemoveFilteredPolicy 从存储器中移除可匹配过滤器的策略规则。
//func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
//	return errors.New("not implemented")
//}
