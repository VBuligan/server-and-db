package storage

import (
	"fmt"
	"github.com/VBuligan/server-and-db/internal/app/models"
	"log"
)

// ArticleRepository * instance of repository (model interface)
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUE ($1,$2,$3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {

	return nil, nil
}

func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	articles, err := ar.SelectAll()

	var founded bool
	if err != nil {
		return nil, founded, err
	}

	var articleFound *models.Article

	for _, a := range articles {
		if a.ID == id {
			articleFound = a
			founded = true
			break
		}
	}

	return articleFound, founded, nil
}

func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// * Куда будем читать
	articles := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}
