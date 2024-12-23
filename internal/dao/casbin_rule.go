package dao

import (
	"github.com/morris-nan/gf-adapter/internal/dao/internal"
)

type casbinRuleDao struct {
	*internal.CasbinRuleDao
}

var (
	CasbinRule = casbinRuleDao{
		internal.NewCasbinRuleDao(),
	}
)
