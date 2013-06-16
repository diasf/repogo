package models

import (
	"database/sql"
	"github.com/robfig/revel"
	db "github.com/robfig/revel/modules/db/app"
)

func init() {
	revel.OnAppStart(func() {
		revel.INFO.Println("Initializing DAOs...")
		var maxConn int
		var sqlConn *sql.DB
		var found bool
		var err error

		revel.INFO.Println("Connecting to the Database")
		db.Init()
		sqlConn = db.Db
		if maxConn, found = revel.Config.Int("db.maxIdleConnections"); !found {
			maxConn = 20
			revel.WARN.Println("Using default value for db.maxIdleConnections", maxConn)
		}
		sqlConn.SetMaxIdleConns(maxConn)

		if err = initUserDao(db.Db); err != nil {
			revel.ERROR.Panicln("Unable to initialize User DAO..", err)
		}

		revel.INFO.Println("DAOs initialized")
	})
}
