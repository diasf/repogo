package models

import (
	"github.com/robfig/revel"
	"database/sql"
)

type User struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
}

type UserDAO interface {
	GetAll() ([]User, error)
	Get(id int) (User, error)
}

func GetUserDAO() UserDAO {
	return userDaoInstance
}

type userDAOImpl struct {
	sqlConn        *sql.DB
	userSelectByIdStmt *sql.Stmt
	userSelectStmt *sql.Stmt
	userUpdateStmt *sql.Stmt
	userInsertStmt *sql.Stmt
	userNexIdStmt  *sql.Stmt
}

var userDaoInstance *userDAOImpl


func initUserDao(sqlConn *sql.DB) (err error) {
	userDaoInstance = &userDAOImpl{}
	userDaoInstance.sqlConn = sqlConn

	revel.TRACE.Println("Creating user Table")
	userDaoInstance.sqlConn.Exec("CREATE TABLE USER (id INTEGER NOT NULL PRIMARY KEY, login TEXT, password TEXT)")

	if err = userDaoInstance.prepareStatements(); err != nil {
		return
	}


	return
}

func (dao *userDAOImpl) prepareStatements() (err error) {
	revel.TRACE.Println("Initializing prepared statements")
	if dao.userSelectByIdStmt,err = dao.sqlConn.Prepare("SELECT id, login, password FROM USER WHERE id=?"); err != nil {
		return
	}
	if dao.userSelectStmt,err = dao.sqlConn.Prepare("SELECT id, login, password FROM USER"); err != nil {
		return
	}
	if dao.userNexIdStmt,err = dao.sqlConn.Prepare("SELECT MAX(id) + 1 FROM USER"); err != nil {
		return
	}
	if dao.userUpdateStmt,err = dao.sqlConn.Prepare("UPDATE USER SET login = ?, password = ? WHERE id = ?"); err != nil {
		return
	}
	if dao.userInsertStmt,err = dao.sqlConn.Prepare("INSERT INTO USER (id, login, password) VALUES (?,?,?)"); err != nil {
		return
	}
	return
}

func (dao *userDAOImpl) nextId() (id int, err error) {
	err = dao.userNexIdStmt.QueryRow().Scan(&id)
	return
}

func (dao *userDAOImpl) GetAll() (users []User, err error) {
	var rs *sql.Rows
	if rs, err = dao.userSelectStmt.Query(); err != nil {
		return
	}
	tmp := []User{}
	for rs.Next() {
		user := User{}
		if err = rs.Scan(&user.Id, &user.Login, &user.Password); err != nil {
			return
		}
		tmp = append(tmp, user)
	}
	users = tmp
	return
}

func (dao *userDAOImpl) Get(id int) (user User, err error) {
	err = dao.userSelectByIdStmt.QueryRow(id).Scan(&user.Id, &user.Login, &user.Password)
	return
}

func (u *User) Save() (err error) {
	if u.Id != 0 {
		revel.TRACE.Println("Updating existing user", u)
		if _, err = userDaoInstance.userUpdateStmt.Exec(u.Login, u.Password, u.Id); err != nil {
			return
		}
	} else {
		var id int
		revel.TRACE.Println("Inserting new user", u)
		if id, err = userDaoInstance.nextId(); err != nil {
			return
		}
		revel.TRACE.Println("New user id:", id)
		if _, err = userDaoInstance.userInsertStmt.Exec(id, u.Login, u.Password); err != nil {
			return
		}
		u.Id = id;
	}
	return
}
