package conn

import(
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const(
	user = "**************"
	pword = "*************"
	host = "*************"
	port = 0000000
	dbname = "lt_stats"
)

func Conn() (*sql.DB) {
	conInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, pword, host, port, dbname)
	db, err := sql.Open("mysql", conInfo)
	if err != nil {
		panic(err.Error())
	}
	// db, err := sql.Open("mysql", "keeper:toor@tcp(localhost:3306)/crypt?parseTime=true")
	// if err != nil {
	// 	panic(err.Error())
	// }

	return db
}