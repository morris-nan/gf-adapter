package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type CasbinRuleDao struct {
	table   string
	group   string
	columns CasbinRuleColumns
}

type CasbinRuleColumns struct {
	Ptype string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}

var casbinRuleColumns = CasbinRuleColumns{
	Ptype: "ptype",
	V0:    "v0",
	V1:    "v1",
	V2:    "v2",
	V3:    "v3",
	V4:    "v4",
	V5:    "v5",
}

func NewCasbinRuleDao() *CasbinRuleDao {
	return &CasbinRuleDao{
		group:   "default",
		table:   "casbin_rule",
		columns: casbinRuleColumns,
	}
}

func (dao *CasbinRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *CasbinRuleDao) Table() string {
	return dao.table
}

func (dao *CasbinRuleDao) Columns() CasbinRuleColumns {
	return dao.columns
}

func (dao *CasbinRuleDao) Group() string {
	return dao.group
}

func (dao *CasbinRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

func (dao *CasbinRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
