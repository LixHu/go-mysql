package sqlDrvier

import (
	"database/sql"
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

// 初始化
func (d Driver) SetDriver(port string, host string, user string, pass string, database string, _type string) (dri *Driver) {
	return &Driver{Port: port, Host: host, User: user, Pass: pass, Database: database, Type: _type}
}

func (d Driver) GetDriver() (dri Driver) {
	return d
}

func (d Driver) Connection() (err error) {
	dsn := d.User + ":" + d.Pass + "@" + d.Type + "(" + d.Host + ":" + d.Port + ")/" + d.Database
	db, err := sql.Open("mysql", dsn)
	db.Ping()
	return err
}

//初始化连接
//func (connection *sqlDriver) Connection() (err error) {
//
//	dsn := connection.user + ":" + connection.pass + "@" + connection._type + "(" + connection.host + ":" + connection.port+ ")/" + connection.database
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func main() {
//
//}
//// 测试连接
//func TestConnection() {
//	sqlDriver := &sqlDriver{"3306", "127.0.0.1", "root", "test", "test", "tcp"}
//	err := connection(sqlDriver)
//	if err != nil {
//		fmt.Print("connetion db failed, err: %v\n", err)
//		return
//	}
//}
