package machines

import (
	"github.com/Extremus-io/gopy/db"
	"github.com/Extremus-io/gopy/log"
	"database/sql"
)

const (
	sql_create_table = `
	CREATE TABLE IF NOT EXISTS machine (
		id 		INT PRIMARY KEY NOT NULL,
		hostname	VARCHAR(30) UNIQUE NOT NULL ,
		extra 		TEXT,
		group 		VARCHAR(200),
		connect_at 	DATETIME NOT NULL DEFAULT DATETIME('now')
	)`
	insert_sql = `INSERT INTO machine (id,hostname,extra,group,connect_at) VALUES (?,?,?,?,?)`
	select_by_id_sql = `SELECT id,hostname,extra,group,connect_at FROM machine where id=?`
)

var machine_ins *sql.Stmt
var machine_sel_by_id *sql.Stmt

func init() {
	_, err := db.DB.Exec(sql_create_table)
	if err != nil {
		panic(err)
	}
	if machine_ins, err = db.DB.Prepare(insert_sql); err != nil {
		panic(err)
	}
	if machine_sel_by_id, err = db.DB.Prepare(select_by_id_sql); err != nil {
		panic(err)
	}
	log.Verbose("Machine DB init complete")
}