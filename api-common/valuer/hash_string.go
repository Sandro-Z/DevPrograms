package valuer

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"git.ana/xjtuana/api-common/crypto/argon2"
)

type HashString string

// GormValue implements the gorm Valuer interface.
func (n HashString) GormValue(ctx context.Context, _ *gorm.DB) clause.Expr {
	hasher := argon2.Hasher{}
	if len(n) != 0 {
		p, _ := hasher.Hash(ctx, []byte(n))
		return clause.Expr{SQL: "?", Vars: []interface{}{string(p)}}
	}
	return clause.Expr{SQL: "?", Vars: []interface{}{""}}
}

// Scan implements the sql.Scanner interface
func (n *HashString) Scan(v interface{}) error {
	var s sql.NullString
	if err := s.Scan(v); err != nil {
		return err
	}
	*n = HashString(s.String)
	return nil
}

func (n *HashString) Compare(s string) error {
	hasher := argon2.Hasher{}
	return hasher.Compare(nil, []byte(*n), []byte(s))
}
