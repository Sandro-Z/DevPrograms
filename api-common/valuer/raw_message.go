package valuer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JSON json.RawMessage

func (n JSON) GormDataType() string {
	return "LONGTEXT"
}

// GormValue implements the gorm Valuer interface.
func (n JSON) GormValue(_ context.Context, _ *gorm.DB) clause.Expr {
	if len(n) == 0 {
		return clause.Expr{SQL: "?", Vars: []interface{}{"null"}}
	}
	return clause.Expr{SQL: "?", Vars: []interface{}{string(n)}}
}

// Scan implements the sql.Scanner interface
func (n *JSON) Scan(value interface{}) error {
	*n = []byte(fmt.Sprintf("%s", value))
	return nil
}

func (n JSON) MarshalJSON() ([]byte, error) {
	if len(n) == 0 {
		return []byte("null"), nil
	}
	return n, nil
}

func (n *JSON) UnmarshalJSON(data []byte) error {
	if n == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*n = append((*n)[0:0], data...)
	return nil
}
