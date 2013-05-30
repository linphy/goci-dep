package goci

import (
	"database/sql"
	"database/sql/driver"
	"github.com/egravert/goci/native"
	"strings"
)

func init() {
	sql.Register("goci", &drv{})
}

type drv struct{}


type transaction struct {
  conn *connection
}

func (d *drv) Open(dsn string) (driver.Conn, error) {

	env, err := native.CreateEnvironment()
	if err != nil {
		return nil, err
	}

	user, pwd, host := parseDsn(dsn)
	svr, err := native.BasicLogin(env, user, pwd, host)
	if err != nil {
		return nil, err
	}

	return &connection{env, svr}, nil
}


// expect the dsn in the format of: user/pwd@host:port/SID
func parseDsn(dsn string) (user, pwd, host string) {
	tokens := strings.SplitN(dsn, "@", 2)
	if len(tokens) > 1 {
		host = tokens[1]
	}
	userpass := strings.SplitN(tokens[0], "/", 2)
	if len(userpass) > 1 {
		pwd = userpass[1]
	}
	user = userpass[0]
	return
}
