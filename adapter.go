package gfadapter

import (
	"context"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/morris-nan/gf-adapter/internal/dao"
	"github.com/morris-nan/gf-adapter/model/entity"
)

func NewAdapter(ctx context.Context) *Adapter {
	return &Adapter{ctx: ctx}
}

type Adapter struct {
	ctx context.Context
}

func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) (err error) {
	lines := make([]*entity.CasbinRule, len(rules))
	for k, rule := range rules {
		lines[k] = savePolicyLine(ptype, rule)
	}
	_, err = dao.CasbinRule.Ctx(a.ctx).Data(lines).Insert()
	return
}

func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) (err error) {
	for _, rule := range rules {
		err = a.RemovePolicy(sec, ptype, rule)
		if err != nil {
			return err
		}
	}
	return
}

func loadPolicyLine(line *entity.CasbinRule, model model.Model) error {
	lineText := line.Ptype
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}
	return persist.LoadPolicyLine(lineText, model)
}

func (a *Adapter) LoadPolicy(model model.Model) (err error) {
	var lines []*entity.CasbinRule
	if err = dao.CasbinRule.Ctx(a.ctx).Scan(&lines); err != nil {
		return err
	}
	for _, line := range lines {
		err = loadPolicyLine(line, model)
		if err != nil {
			return
		}
	}
	return
}

func savePolicyLine(ptype string, rule []string) *entity.CasbinRule {
	line := &entity.CasbinRule{}
	line.Ptype = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}
	return line
}

func (a *Adapter) SavePolicy(model model.Model) (err error) {

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			_, err = dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
			if err != nil {
				return
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			_, err = dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
			if err != nil {
				return
			}
		}
	}
	return
}

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) (err error) {
	line := savePolicyLine(ptype, rule)
	_, err = dao.CasbinRule.Ctx(a.ctx).Data(line).Insert()
	return
}

func (a *Adapter) rawDelete(line *entity.CasbinRule) (err error) {
	db := dao.CasbinRule.Ctx(a.ctx).Where("ptype = ?", line.Ptype)
	if line.V0 != "" {
		db = db.Where("v0 = ?", line.V0)
	}
	if line.V1 != "" {
		db = db.Where("v1 = ?", line.V1)
	}
	if line.V2 != "" {
		db = db.Where("v2 = ?", line.V2)
	}
	if line.V3 != "" {
		db = db.Where("v3 = ?", line.V3)
	}
	if line.V4 != "" {
		db = db.Where("v4 = ?", line.V4)
	}
	if line.V5 != "" {
		db = db.Where("v5 = ?", line.V5)
	}
	_, err = db.Delete()
	return
}

func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	return a.rawDelete(line)
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := &entity.CasbinRule{}
	line.Ptype = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	return a.rawDelete(line)
}
