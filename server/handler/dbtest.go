package handler

import (
	"fmt"
	"net/http"

	"github.com/gumpen/comp/server/infrastructure"
)

func Dbtest(w http.ResponseWriter, r *http.Request) {
	db, err := infrastructure.DbConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var name string
		var sex string

		if err := rows.Scan(&name, &sex); err != nil {
			panic(err)
		}
		fmt.Fprintln(w, name)
		fmt.Fprintln(w, sex)
	}

}
