package do

import "github.com/gogf/gf/v2/frame/g"

type CasbinRule struct {
	g.Meta `orm:"table:casbin_rule, do:true"`
	Ptype  interface{} //
	V0     interface{} //
	V1     interface{} //
	V2     interface{} //
	V3     interface{} //
	V4     interface{} //
	V5     interface{} //
}
