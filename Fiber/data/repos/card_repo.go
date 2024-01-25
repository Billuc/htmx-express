package repos

import (
	"context"
	"database/sql"
	"fax/data/db"
	"fax/data/models"

	"github.com/google/uuid"
)

type CardRepo struct {
	db *sql.DB
}

func (repo CardRepo) GetManyCards() ([]models.Card, error) {
	rows, err := repo.db.Query("SELECT * FROM cards")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []models.Card{}
	for rows.Next() {
		card := models.Card{}
		err := rows.Scan(&card.ID, &card.Title, &card.Description, &card.HasDate, &card.Date, &card.Frequency, &card.FrequencyUnit, &card.TagID, &card.ParentID)

		if err != nil {
			return nil, err
		}

		data = append(data, card)
	}

	return data, nil
}

func (repo CardRepo) GetCard(id string) (*models.Card, error) {
	row := repo.db.QueryRow("SELECT * FROM cards WHERE id = ?", id)

	card := models.Card{}
	err := row.Scan(&card.ID, &card.Title, &card.Description, &card.HasDate, &card.Date, &card.Frequency, &card.FrequencyUnit, &card.TagID, &card.ParentID)

	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (repo CardRepo) CreateCard(title string) (*models.Card, error) {
	tx, err := repo.db.BeginTx(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	var id string = uuid.NewString()
	_, err = tx.Exec("INSERT INTO cards (id, title, has_date) VALUES (?, ?, ?)", id, title, false)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return repo.GetCard(id)
}

func GetCardRepo() (*CardRepo, error) {
	db, err := db.GetDB()

	if err != nil {
		return nil, err
	}

	return &CardRepo{db}, nil
}
