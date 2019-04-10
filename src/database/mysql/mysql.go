package mysql

import (
	"orderfood/src/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlDb struct {
	Connect func() (*sqlx.DB, error)
}

func NewDb(dbCfg config.DbConfig) (*mysqlDb, error) {
	d := &mysqlDb{Connect: func() (*sqlx.DB, error) {
		db, err := sqlx.Open("mysql", dbCfg.MysqlURL())
		return db, err
	}}

	//check db
	db, err := d.Connect()
	defer db.Close()

	return d, err
}

func Rebuild(dbCfg config.DbConfig) error {
	d, err := NewDb(dbCfg)
	if err != nil {
		return err
	}

	//check db struct
	db, err := d.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	dropSQL := `DROP TABLE user_info;`
	_, err = db.Exec(dropSQL)

	createSQL := `CREATE TABLE user_info (
	id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(45) NOT NULL,
	password VARCHAR(45) NULL,
	PRIMARY KEY (id),
	UNIQUE INDEX username_UNIQUE (username ASC) VISIBLE);`

	_, err = db.Exec(createSQL)

	return err
}
