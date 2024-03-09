package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/themuks/blog-api/app/queries"
)

type Queries struct {
	*queries.ArticleQueries
}

func OpenDBConnection() (*Queries, error) {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	client, err := sql.Open("mysql", url)

	if err != nil {
		return nil, err
	}

	return &Queries{
		ArticleQueries: &queries.ArticleQueries{DB: client},
	}, nil
}
