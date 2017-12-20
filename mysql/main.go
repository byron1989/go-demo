// mysql project main.go
package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
	db   *sql.DB
)

func initDb(dataSourceName string, maxOpenConns int, maxIdleConns int) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("open database error, ", err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	err = db.Ping()
	if err != nil {
		fmt.Println("db ping error, ", err.Error())
		return nil, err
	}

	fmt.Println("db init success ")
	return db, nil
}

func closeDb(db *sql.DB) {
	if db != nil {
		db.Close()
		fmt.Println("db closed")
	}
}

func testInsert(db *sql.DB) error {
	sql := "insert into test(value, status) values (?,?)"
	stmt, err := db.Prepare(sql)
	if nil != err {
		return err
	}

	defer stmt.Close()
	for i := 0; i < 5; i++ {
		rs, err := stmt.Exec(strconv.Itoa(i), "valid")
		if nil != err {
			return err
		}

		lastid, err := rs.LastInsertId()
		if nil != err {
			return err
		}

		affected, err := rs.RowsAffected()
		if nil != err {
			return err
		}

		fmt.Println("test insert ", lastid, affected)
	}

	return nil
}

func testUpdate(db *sql.DB) error {
	sql := "update test set status = ? where value = ?"
	stmt, err := db.Prepare(sql)
	if nil != err {
		return err
	}

	defer stmt.Close()

	rs, err := stmt.Exec("invalid", 3)
	if nil != err {
		return err
	}

	lastid, err := rs.LastInsertId()
	if nil != err {
		return err
	}

	affected, err := rs.RowsAffected()
	if nil != err {
		return err
	}

	fmt.Println("test update ", lastid, affected)

	return nil
}

func testDelete(db *sql.DB) error {
	sql := "delete from  test where value = ?"
	stmt, err := db.Prepare(sql)
	if nil != err {
		return err
	}

	defer stmt.Close()

	rs, err := stmt.Exec(2)
	if nil != err {
		return err
	}

	lastid, err := rs.LastInsertId()
	if nil != err {
		return err
	}

	affected, err := rs.RowsAffected()
	if nil != err {
		return err
	}

	fmt.Println("test delete ", lastid, affected)

	return nil
}

func testQueryRow(db *sql.DB) error {
	var id int64
	var value string
	var status string
	var updateAt string
	var createAt string

	sql := "select * from test"
	err := db.QueryRow(sql).Scan(&id, &value, &status, &updateAt, &createAt)
	if nil != err {
		return err
	}

	fmt.Println("test query row:", id, value, status, updateAt, createAt)

	return nil
}

func testQuery(db *sql.DB) error {
	var id int64
	var value string
	var status string
	var updateAt string
	var createAt string

	sql := "select * from test where id < ?"
	rows, err := db.Query(sql, 100)
	if nil != err {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&id, &value, &status, &updateAt, &createAt)
		if nil != err {
			return err
		}
		fmt.Println("test query:", id, value, status, updateAt, createAt)
	}

	return nil
}

func main() {
	dataSourceString := "root:root123@tcp(127.0.0.1:3306)/ares?charset=utf8"
	db, err := initDb(dataSourceString, 5, 5)
	if nil != err {
		fmt.Println(err.Error())
		return
	}

	if err := testInsert(db); nil != err {
		fmt.Println(err.Error())
		return
	}

	if err := testUpdate(db); nil != err {
		fmt.Println(err.Error())
		return
	}

	if err := testDelete(db); nil != err {
		fmt.Println(err.Error())
		return
	}

	if err := testQueryRow(db); nil != err {
		fmt.Println(err.Error())
		return
	}

	if err := testQuery(db); nil != err {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("hello world.")
}
