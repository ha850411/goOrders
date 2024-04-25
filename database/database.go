package database

import (
	"database/sql"
	"fmt"
	"goOrders/conf"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbInfo struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
}

var dbInfo = DbInfo{
	Host:     conf.Settings.Database.Host,
	Port:     conf.Settings.Database.Port,
	Username: conf.Settings.Database.Username,
	Password: conf.Settings.Database.Password,
	Dbname:   conf.Settings.Database.Dbname,
}

func DbConnect() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbInfo.Username, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.Dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func GormConnect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=False&loc=Local", dbInfo.Username, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func QueryToMap(db *sql.DB, sSql string) ([]map[string]interface{}, error) {

	rows, err := db.Query(sSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	count := len(columns)
	values, valuesPoint := make([]interface{}, count), make([]interface{}, count)

	mData := make([]map[string]interface{}, 0)
	for rows.Next() {

		for i := 0; i < count; i++ {
			valuesPoint[i] = &values[i]
		}

		rows.Scan(valuesPoint...)

		entry := make(map[string]interface{})

		for i, val := range values {
			key := columns[i]
			var v interface{}
			b, ok := val.([]byte)
			if ok {
				// 字符切片转为字符串
				v = string(b)
			} else {
				v = val
			}
			entry[key] = v
		}
		mData = append(mData, entry)
	}
	return mData, nil
}
