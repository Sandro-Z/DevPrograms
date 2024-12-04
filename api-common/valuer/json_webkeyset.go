package valuer

import (
	"context"
	"encoding/json"
	"fmt"

	"gopkg.in/square/go-jose.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JSONWebKeySet struct {
	*jose.JSONWebKeySet
}

func (n JSONWebKeySet) GormDataType() string {
	return "LONGTEXT"
}

// GormValue implements the gorm Valuer interface.
func (n *JSONWebKeySet) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	s, err := json.Marshal(n)
	if err != nil {
		return clause.Expr{SQL: "?", Vars: []interface{}{""}}
	}
	return clause.Expr{SQL: "?", Vars: []interface{}{s}}
}

// Scan implements the sql.Scanner interface
func (n *JSONWebKeySet) Scan(v interface{}) error {
	s := fmt.Sprintf("%s", v)
	if len(s) == 0 {
		return nil
	}
	return json.Unmarshal([]byte(s), n)
}

func (n *JSONWebKeySet) ToJson() (string, error) {
	s, err := json.Marshal(n)
	return string(s), err
}
