package controllers

import (
	"github.com/revel/revel"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SelectQuery struct {
	*revel.Controller
}

func (c SelectQuery) Index() revel.Result {
	db, err := sql.Open("mysql", "root:@/revel")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

//	SELECT
	rows, err := db.Query("SELECT id, title FROM test")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	fmt.Println(rows)

	type Row struct {
		Id int
		Title string
	}

	results := []Row{}
	for rows.Next() {
//		代入用変数を定義
		var Id int
		var Title string

//		Query値を変数へ代入
		rows.Scan(&Id, &Title)

//		Example 1
		results = append(results, Row{ Id, Title })
//		Example 2
//		var row = Row{Id, Title}
//		results = append(results, row)
//		Example 3
//		var row = Row{}
//		row.Id = Id
//		row.Title = Title
//		results = append(results, row)
	}

	return c.Render(results)
}
