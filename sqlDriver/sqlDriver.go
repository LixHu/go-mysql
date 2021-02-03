package sqlDrvier

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Driver struct {
	Port     string
	Host     string
	User     string
	Pass     string
	Database string
	Type     string
}

var db *sql.DB

type Test struct {
	id   int
	test string
}

// 初始化
func (d Driver) SetDriver(port string, host string, user string, pass string, database string, _type string) (dri *Driver) {
	return &Driver{Port: port, Host: host, User: user, Pass: pass, Database: database, Type: _type}
}

func (d Driver) GetDriver() (dri Driver) {
	return d
}

func (d Driver) connection() (err error) {
	dsn := d.User + ":" + d.Pass + "@" + d.Type + "(" + d.Host + ":" + d.Port + ")/" + d.Database
	db, err = sql.Open("mysql", dsn)
	return err
}

func (d Driver) Select() (err error, row sql.Row) {
	err = d.connection()
	if err != nil {
		return
	}
	rows, err := db.Query("select * from test")

	defer db.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var t Test
		err1 := rows.Scan(&t.id, &t.test)
		if err1 != nil {
			return
		}
		fmt.Println(t)
	}
	//for rows.Next() {
	//	row, err := rows.Columns()
	//	fmt.Println(row, err)
	//}
	return err, row
}
