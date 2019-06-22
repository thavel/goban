package auth

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // used by gorm
)

// Anonymous role
const Anonymous = "anonymous"

var (
	conf     = casbin.NewModel()
	enforcer *casbin.Enforcer

	// This cache prevents querying the db everytime a request is made
	memcache  = newCache(100000, time.Minute, time.Minute*3)
	bearer, _ = regexp.Compile(`^Bearer (.*)$`)
)

func init() {
	conf.AddDef("r", "r", `sub, obj, act`)
	conf.AddDef("p", "p", `sub, obj, act`)
	conf.AddDef("e", "e", `some(where (p.eft == allow))`)
	conf.AddDef("m", "m", `r.sub == p.sub && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")`)
}

// Setup init/load the policy backend database.
func Setup(db *gorm.DB) error {
	log.SetOutput(ioutil.Discard)
	defer log.SetOutput(os.Stderr)
	adapter := gormadapter.NewAdapterByDB(db)
	enf, err := casbin.NewEnforcerSafe(conf, adapter)
	if err != nil {
		return err
	}
	enf.EnableLog(false)
	enf.LoadPolicy()
	enforcer = enf
	return nil
}

// Save policies into the db.
func Save() {
	enforcer.SavePolicy()
}

// Enforcer object.
func Enforcer() *casbin.Enforcer {
	return enforcer
}
