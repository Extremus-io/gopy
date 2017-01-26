package auth

import (
	"database/sql"
	"github.com/Extremus-io/gopy/db"
	"github.com/Extremus-io/gopy/log"
)

const tableUsers = "gopy_auth_users"

const colEmail = "email"
const colPassHash = "pass_sha256"
const colIsSuperUser = "is_superuser"
const colIsActive = "is_active"

const sql_db_create_table = `
	CREATE TABLE IF NOT EXISTS ` + tableUsers + ` (
		` + colEmail + ` 	VARCHAR(100) UNIQUE NOT NULL,
		` + colPassHash + `	TEXT NOT NULL,
		` + colIsSuperUser + `	BOOL DEFAULT FALSE,
		` + colIsActive + `	BOOL DEFAULT TRUE
	)
`

const getUserByEmailSQL = `SELECT
	 ` + colEmail + `,
	 ` + colPassHash + `,
	 ` + colIsSuperUser + `,
	 ` + colIsActive + ` FROM ` + tableUsers + ` WHERE email=?`

const insertUserSQL = `INSERT INTO ` + tableUsers + `(
	 ` + colEmail + `,
	 ` + colPassHash + `,
	 ` + colIsSuperUser + `,
	 ` + colIsActive + `) VALUES (?,?,?,?)`

const deleteUserSQL = `DELETE FROM ` + tableUsers + ` WHERE ` + colEmail + `=?`

const updateUserSQL = `UPDATE ` + tableUsers + ` SET
	` + colPassHash + `=?,
	` + colIsSuperUser + `=?,
	` + colIsActive + `=? WHERE ` + colEmail + `=?`

var stmtGetUserByEmail *sql.Stmt
var stmtInsertUser *sql.Stmt
var stmtDeleteUser *sql.Stmt
var stmtUpdateUser *sql.Stmt

func init() {
	// initializing table
	_, err := db.DB.Exec(sql_db_create_table)
	if err != nil {
		panic(err)
	}

	// initializing select statement
	if stmtGetUserByEmail, err = db.DB.Prepare(getUserByEmailSQL); err != nil {
		panic(err)
	}

	// initializing insert statement
	if stmtInsertUser, err = db.DB.Prepare(insertUserSQL); err != nil {
		panic(err)
	}

	// initializing delete statement
	if stmtDeleteUser, err = db.DB.Prepare(deleteUserSQL); err != nil {
		panic(err)
	}

	// initializing update statement
	if stmtUpdateUser, err = db.DB.Prepare(updateUserSQL); err != nil {
		panic(err)
	}
	log.Verbose("Auth database init complete")
}