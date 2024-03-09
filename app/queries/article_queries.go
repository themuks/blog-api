package queries

import (
	"database/sql"

	"github.com/themuks/blog-api/app/models"
)

type ArticleQueries struct {
	*sql.DB
}

func (q *ArticleQueries) GetArticles() ([]models.Article, error) {
	rows, err := q.Query("SELECT id, title, text FROM articles")
	if err != nil {
		return nil, err
	}
	articles := []models.Article{}
	for rows.Next() {
		article := models.Article{}
		err = rows.Scan(&article.ID, &article.Title, &article.Text)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (q *ArticleQueries) GetArticleById(id int) (models.Article, error) {
	row := q.QueryRow("SELECT id, title, text FROM articles WHERE id = ?", id)
	article := models.Article{}
	err := row.Scan(&article.ID, &article.Title, &article.Text)
	if err != nil {
		return article, err
	}
	return article, nil
}
