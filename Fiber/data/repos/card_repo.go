package repos

import (
	"database/sql"
	"fax/data/db"
	"fax/data/models"
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

func GetCardRepo() (*CardRepo, error) {
	db, err := db.GetDB()

	if err != nil {
		return nil, err
	}

	return &CardRepo{db}, nil
}
