package controllers

import (
	"github.com/revel/revel"
	"database/sql"
	"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	db, err := sql.Open("mysql", "root:@/revel")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

//	CREATE TABLE
	query := `
		CREATE TABLE test (
		  id int(11) unsigned NOT NULL AUTO_INCREMENT,
		  title varchar(255) DEFAULT NULL,
		  PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
	`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}

//	INSERT
	for i := 1; i <= 20; i++ {
		title := "test" + strconv.Itoa(i)
		_, err = db.Exec(`INSERT test (title) VALUES (?)`, title)
	}
	if err != nil {
		panic(err.Error())
	}

//	SELECT
	rows, err := db.Query("SELECT id, title FROM test")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	fmt.Println(rows)
	for rows.Next() {
		var ID int
		var post_title string
		if err := rows.Scan(&ID, &post_title); err != nil {
			panic(err.Error())
		}
		fmt.Println(ID, post_title)
	}



  	greeting:="Aloha World"
	return c.Render(greeting)
}
