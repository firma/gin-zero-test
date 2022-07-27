// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"miya/api/internal/model"
)

func newSysMenu(db *gorm.DB) sysMenu {
	_sysMenu := sysMenu{}

	_sysMenu.sysMenuDo.UseDB(db)
	_sysMenu.sysMenuDo.UseModel(&model.SysMenu{})

	tableName := _sysMenu.sysMenuDo.TableName()
	_sysMenu.ALL = field.NewField(tableName, "*")
	_sysMenu.ID = field.NewInt64(tableName, "id")
	_sysMenu.MenuName = field.NewString(tableName, "menu_name")
	_sysMenu.MenuType = field.NewString(tableName, "menu_type")
	_sysMenu.PMenuID = field.NewInt64(tableName, "p_menu_id")
	_sysMenu.MenuIcon = field.NewString(tableName, "menu_icon")
	_sysMenu.Xh = field.NewInt32(tableName, "xh")
	_sysMenu.AuthCode = field.NewString(tableName, "auth_code")
	_sysMenu.Remark = field.NewString(tableName, "remark")
	_sysMenu.Path = field.NewString(tableName, "path")
	_sysMenu.Redirect = field.NewString(tableName, "redirect")
	_sysMenu.CreateTime = field.NewInt64(tableName, "create_time")
	_sysMenu.UpdateTime = field.NewInt64(tableName, "update_time")
	_sysMenu.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysMenu.DeleteTime = field.NewField(tableName, "delete_time")
	_sysMenu.IfLeaf = field.NewString(tableName, "if_leaf")

	_sysMenu.fillFieldMap()

	return _sysMenu
}

type sysMenu struct {
	sysMenuDo sysMenuDo

	ALL        field.Field
	ID         field.Int64
	MenuName   field.String
	MenuType   field.String
	PMenuID    field.Int64
	MenuIcon   field.String
	Xh         field.Int32
	AuthCode   field.String
	Remark     field.String
	Path       field.String
	Redirect   field.String
	CreateTime field.Int64
	UpdateTime field.Int64
	ModifyTime field.Time
	DeleteTime field.Field
	IfLeaf     field.String

	fieldMap map[string]field.Expr
}

func (s sysMenu) Table(newTableName string) *sysMenu {
	s.sysMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysMenu) As(alias string) *sysMenu {
	s.sysMenuDo.DO = *(s.sysMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysMenu) updateTableName(table string) *sysMenu {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt64(table, "id")
	s.MenuName = field.NewString(table, "menu_name")
	s.MenuType = field.NewString(table, "menu_type")
	s.PMenuID = field.NewInt64(table, "p_menu_id")
	s.MenuIcon = field.NewString(table, "menu_icon")
	s.Xh = field.NewInt32(table, "xh")
	s.AuthCode = field.NewString(table, "auth_code")
	s.Remark = field.NewString(table, "remark")
	s.Path = field.NewString(table, "path")
	s.Redirect = field.NewString(table, "redirect")
	s.CreateTime = field.NewInt64(table, "create_time")
	s.UpdateTime = field.NewInt64(table, "update_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.DeleteTime = field.NewField(table, "delete_time")
	s.IfLeaf = field.NewString(table, "if_leaf")

	s.fillFieldMap()

	return s
}

func (s *sysMenu) WithContext(ctx context.Context) *sysMenuDo { return s.sysMenuDo.WithContext(ctx) }

func (s sysMenu) TableName() string { return s.sysMenuDo.TableName() }

func (s sysMenu) Alias() string { return s.sysMenuDo.Alias() }

func (s *sysMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["id"] = s.ID
	s.fieldMap["menu_name"] = s.MenuName
	s.fieldMap["menu_type"] = s.MenuType
	s.fieldMap["p_menu_id"] = s.PMenuID
	s.fieldMap["menu_icon"] = s.MenuIcon
	s.fieldMap["xh"] = s.Xh
	s.fieldMap["auth_code"] = s.AuthCode
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["path"] = s.Path
	s.fieldMap["redirect"] = s.Redirect
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["delete_time"] = s.DeleteTime
	s.fieldMap["if_leaf"] = s.IfLeaf
}

func (s sysMenu) clone(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceDB(db)
	return s
}

type sysMenuDo struct{ gen.DO }

func (s sysMenuDo) Debug() *sysMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysMenuDo) WithContext(ctx context.Context) *sysMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysMenuDo) ReadDB() *sysMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysMenuDo) WriteDB() *sysMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysMenuDo) Clauses(conds ...clause.Expression) *sysMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysMenuDo) Returning(value interface{}, columns ...string) *sysMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysMenuDo) Not(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysMenuDo) Or(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysMenuDo) Select(conds ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysMenuDo) Where(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysMenuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysMenuDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysMenuDo) Order(conds ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysMenuDo) Distinct(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysMenuDo) Omit(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysMenuDo) Join(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysMenuDo) Group(cols ...field.Expr) *sysMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysMenuDo) Having(conds ...gen.Condition) *sysMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysMenuDo) Limit(limit int) *sysMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysMenuDo) Offset(offset int) *sysMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysMenuDo) Unscoped() *sysMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysMenuDo) Create(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysMenuDo) CreateInBatches(values []*model.SysMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysMenuDo) Save(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysMenuDo) First() (*model.SysMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Take() (*model.SysMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Last() (*model.SysMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Find() ([]*model.SysMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysMenu), err
}

func (s sysMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error) {
	buf := make([]*model.SysMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysMenuDo) FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysMenuDo) Attrs(attrs ...field.AssignExpr) *sysMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysMenuDo) Assign(attrs ...field.AssignExpr) *sysMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysMenuDo) Joins(fields ...field.RelationField) *sysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysMenuDo) Preload(fields ...field.RelationField) *sysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysMenuDo) FirstOrInit() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FirstOrCreate() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s *sysMenuDo) withDO(do gen.Dao) *sysMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
