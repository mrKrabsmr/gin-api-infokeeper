package dao

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mrKrabsmr/infokeeper/internal/app/models"
)

type InfoDAO struct {
	*sqlx.DB
}

func NewInfoDAO(db *sqlx.DB) *InfoDAO {
	return &InfoDAO{db}
}

func (dao *InfoDAO) GetObjByID(id uuid.UUID) (models.Info, error) {
	var info models.Info

	query := `SELECT * FROM info WHERE id = $1`

	if err := dao.Get(&info, query, id); err != nil {
		return info, err
	}

	return info, nil
}

func (dao *InfoDAO) GetObjByKey(key []byte) (models.Info, error) {
	var info models.Info

	query := `SELECT * FROM info WHERE key = $1`

	if err := dao.Get(&info, query, key); err != nil {
		return info, err
	}

	return info, nil
}

func (dao *InfoDAO) GetIPAddressByKey(key []byte) (string, error) {
	var ip string

	query := `SELECT c.ip_address FROM info i JOIN clients c ON i.client_id = c.id WHERE key = $1`

	if err := dao.Get(&ip, query, key); err != nil {
		return "", err
	}

	return ip, nil
}

func (dao *InfoDAO) GetCountValuesByIPAddress(ip string) (int, error) {
	var count int

	query := `SELECT COUNT(i.*) FROM info i 
			  JOIN clients c ON i.client_id = c.id 
			  WHERE c.ip_address = $1`

	if err := dao.Get(&count, query, ip); err != nil {
		return -1, err
	}

	return count, nil
}

func (dao *InfoDAO) Create(info *models.Info) (uuid.UUID, error) {
	var infoID uuid.UUID

	query := `INSERT INTO info(id, key, value, client_id, read_only) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	if err := dao.QueryRow(
		query, info.ID, info.Key, info.Value, info.ClientID, info.ReadOnly,
	).Scan(&infoID); err != nil {
		return infoID, err
	}

	return infoID, nil
}

func (dao *InfoDAO) Update(info *models.Info) error {
	query := `UPDATE info SET value = $1, read_only = $2 WHERE id = $3`

	if _, err := dao.Exec(query, info.Value, info.ReadOnly, info.ID); err != nil {
		return err
	}

	return nil
}

func (dao *InfoDAO) Delete(key []byte) error {
	query := `DELETE FROM info WHERE key = $1`

	if _, err := dao.Exec(query, key); err != nil {
		return err
	}

	return nil
}
