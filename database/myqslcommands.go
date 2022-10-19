package database

import (
	"database/sql"
	"fmt"
	"log"
)

type SQLCommand struct {
	SqlConn *MySQLDB
}

func (sqlCommand SQLCommand) Ping() error {

	db, err := sql.Open(sqlCommand.SqlConn.Driver, sqlCommand.SqlConn.DataSource)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("FAIL TO CONNECT TO SQL SERVER!")
		return err
	}

	fmt.Println("Connected!")
	return nil
}

func (sqlCommand SQLCommand) ExecuteSQLCommand(command string) {

	db, err := sql.Open(sqlCommand.SqlConn.Driver, sqlCommand.SqlConn.DataSource)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query, err := db.Query(command)
	if err != nil {
		panic(err.Error())
	}

	defer query.Close()
}
