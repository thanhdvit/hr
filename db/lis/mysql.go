// Package db Tương tác với LIS database
//
// @Author	Đào Văn Thanh
// @Date		2019-05-31
//
package db

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//MySQLLis Lis database
func MySQLLis() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "lis"
	dbMySqlUser := os.Getenv("db_lis_user")
	dbMySqlPass := os.Getenv("db_lis_pwd")
	domainMySql := os.Getenv("db_lis_ip")

	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbMySqlUser, dbMySqlPass, domainMySql, dbName)
	db, err = sql.Open(dbDriver, conn)

	return db, err
}

//ValidUser kiểm tra thông tin user trên LIS
func ValidUser(user, pwd string) (userID int64, err error) {
	db, err := MySQLLis()
	if err != nil {
		return 0, errors.New("Không kết nối với database")
	}
	defer db.Close()

	matkhau := strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(pwd))))

	sqlString := "SELECT employee_id FROM employees WHERE user_name=? AND password=?"
	err = db.QueryRow(sqlString, user, matkhau).Scan(&userID)
	if err != nil {
		return 0, err
	}

	if userID == 0 {
		return 0, errors.New("username hoặc password không hợp lệ")
	}
	return userID, nil
}
