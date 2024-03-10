package database

import (
	"database/sql"
	"fmt"
	"github.com/Mart1n-Eden/web-diary/internal/model"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	// host = "host.docker.internal"
	host     = "localhost"
	// host     = "0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func Init() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	println("Sucsess connection to databse")
}

func Close() { defer db.Close() }

func GetTemplate(num int) []model.Article {
	rows, err := db.Query("SELECT * FROM articles ORDER BY id desc")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for ; num != 0; num-- {
		rows.Next()
	}

	result := make([]model.Article, 3)

	for i := 0; i < 3 && rows.Next(); i++ {
		if err := rows.Scan(&result[i].ID, &result[i].Title, &result[i].Body); err != nil {
			panic(err)
		}
	}

	return result
}

func CountArticles() int {
	rows, err := db.Query("SELECT COUNT(*) FROM articles")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result int
	rows.Next()
	if err := rows.Scan(&result); err != nil {
		panic(err)
	}

	return result
}

func TakeArticle(a *model.Article, i int) {
	textQuery := fmt.Sprintf("SELECT * FROM articles WHERE id = %d;", i)

	rows, err := db.Query(textQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&a.ID, &a.Title, &a.Body); err != nil {
			panic(err)
		}
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}

func SaveArticle(a model.Article) {
	_, err := db.Exec("INSERT INTO articles (title, body) VALUES ($1, $2)", a.Title, a.Body)
	if err != nil {
		panic(err)
	}
}
