package goci

import (
	"database/sql/driver"
	"github.com/egravert/goci/native"
)

type connection struct {
	env *native.EnvHandle
	svr *native.SvrHandle
}

func (conn *connection) Begin() (driver.Tx, error) {
  stmt, err := native.Prepare(conn.env, "BEGIN")
  if err != nil {
    return nil, err
  }

  if err = native.Exec(conn.env, conn.svr, stmt); err != nil {
    return nil, err
  }
  return &transaction{conn}, nil
}

func (conn *connection) Prepare(query string) (driver.Stmt, error) {
	stmt, err := native.Prepare(conn.env, query)
	if err != nil {
		return nil, err
	}
	return &statement{handle: stmt, conn: conn}, nil
}

func (conn *connection) Close() error {
	return nil
}
