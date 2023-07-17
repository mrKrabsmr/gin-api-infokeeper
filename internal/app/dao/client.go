package dao

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mrKrabsmr/infokeeper/internal/app/models"
)

type ClientDAO struct {
	*sqlx.DB
}

func NewClientDAO(db *sqlx.DB) *ClientDAO {
	return &ClientDAO{db}
}

func (dao *ClientDAO) Create(client *models.Client) (uuid.UUID, error) {
	exists, err := dao.CheckExists(client.IPAddress)

	if err != nil {
		return client.ID, err
	}

	if !exists {
		query := `INSERT INTO clients(id, ip_address) VALUES($1, $2)`

		if _, err := dao.Exec(query, client.ID, client.IPAddress); err != nil {
			return client.ID, err
		}

		return client.ID, nil
	}

	clientObj, err := dao.GetByIP(client.IPAddress)
	if err != nil {
		return uuid.UUID{}, err
	}

	return clientObj.ID, nil

}

func (dao *ClientDAO) CheckExists(ip string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT * FROM clients WHERE ip_address = $1)`

	if err := dao.Get(&exists, query, ip); err != nil {
		return false, err
	}

	return exists, nil
}

func (dao *ClientDAO) GetByIP(ip string) (models.Client, error) {
	var client models.Client

	query := `SELECT * FROM clients WHERE ip_address = $1`

	if err := dao.Get(&client, query, ip); err != nil {
		return client, err
	}

	return client, nil
}
