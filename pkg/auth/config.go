package auth

import (
	"fmt"
	"io/ioutil"
	bilog "log"
	"os"
	"regexp"
	"time"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/mysql" // used by gorm

	"github.com/thavel/goban/pkg/database"
)

const Anonymous = "anonymous"

var (
	enforcer  *casbin.Enforcer
	bearer, _ = regexp.Compile(`^Bearer (.*)$`)

	// This cache prevents querying the db everytime a request is made
	memcache = newCache(100000, time.Minute, time.Minute*3)
)

// SetupPolicies init/load the policy backend database.
func SetupPolicies(config database.Config) error {
	bilog.SetOutput(ioutil.Discard)
	defer bilog.SetOutput(os.Stderr)
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database,
	)
	a := gormadapter.NewAdapter("mysql", uri, true)
	e, err := casbin.NewEnforcerSafe("./rbac.conf", a)
	if err != nil {
		return err
	}
	e.EnableLog(false)
	e.LoadPolicy()
	enforcer = e
	return nil
}

// SavePolicies saves policies into the db.
func SavePolicies() {
	enforcer.SavePolicy()
}

// Enforcer object.
func Enforcer() *casbin.Enforcer {
	return enforcer
}
