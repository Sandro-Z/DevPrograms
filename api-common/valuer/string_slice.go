package valuer

import (
	"context"
	"database/sql"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StringSlice []string

func (n StringSlice) GormDataType() string {
	return "LONGTEXT"
}

// GormValue implements the gorm Valuer interface.
func (n StringSlice) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	if len(n) == 0 {
		return clause.Expr{SQL: "?", Vars: []interface{}{""}}
	}
	return clause.Expr{SQL: "?", Vars: []interface{}{strings.Join(n, "|")}}
}

// Scan implements the sql.Scanner interface
func (n *StringSlice) Scan(v interface{}) error {
	var s sql.NullString
	if err := s.Scan(v); err != nil {
		return err
	}
	if len(s.String) > 0 {
		*n = strings.Split(string(s.String), "|")
	}
	return nil
}
