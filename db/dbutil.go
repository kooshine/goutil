package db

import (
    "fmt"
    "strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var SQLManager *SQLManage

type SQLManage struct {
    lemon *sql.DB
}
func init() {
    SQLManager = &SQLManage {
        lemon: nil,
    }
}

func (s *SQLManage) DBConnect(dbUrl string) {
    var err error
	s.lemon, err = sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
    fmt.Println("Lemon DB connect OK")
}
func (s *SQLManage) Count(tablename string) int64 {
    /*
        查询符合条件的记录数
    */
    count_lang := fmt.Sprintf("SELECT count(1) FROM %s", tablename)
    rows, err := s.lemon.Query(count_lang)
    defer rows.Close()
    if err != nil {
        panic(err)
        return -1
    }
    var total int64
    for rows.Next() {
        rows.Scan(&total)
    }
    return total
}
func (s *SQLManage) Search(tablename string, fields string, indexMap string) []map[string]string {
    /*
        查找数据
    */
    lang := fmt.Sprintf("SELECT %s FROM %s WHERE %s", fields, tablename, indexMap)

    var result []map[string]string
    rows, err := s.lemon.Query(lang)
    if err != nil {
        panic(err)
        return result
    }

    columns, _ := rows.Columns()

    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        res := make(map[string]string)
        rows.Scan(scanArgs...)
        for i, col := range values {
            res[columns[i]] = string(col)
        }
        result = append(result, res)
    }

    return result
}

func (s *SQLManage) Insert(tablename string, data map[string]string) int64 {
    /*
    lang := "INSERT INTO packet(topic, channel, message) VALUES('test5','abcde', '123456')"
    fmt.Println(lang)

    stat, err := s.lemon.Prepare(lang)
    if err != nil {
        panic(err)
        return -1
    }
    */
    var key string
    var value string
    for item, val := range data {
        key = strings.Join([]string{key, item}, ", ")
        value = strings.Join([]string{value, val}, "','")
    }
    fmt.Println(key[1:], "\n", value[2:])
    lang := "INSERT INTO " + tablename + "(" + key[1:] +")" + "VALUES(" + value[2:]+"')"
    fmt.Println(lang)

    ret, err := s.lemon.Exec(lang)
    if err != nil {
        panic(err)
        return -1
    }
    id , err := ret.LastInsertId()
    if err != nil {
        panic(err)
        return -1
    }
    return id
}
