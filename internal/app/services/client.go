package services

import (
	"github.com/google/uuid"
	"github.com/mrKrabsmr/infokeeper/internal/app/dao"
	"github.com/mrKrabsmr/infokeeper/internal/app/models"
)

type ClientService struct {
	DAO *dao.ClientDAO
}

func NewClientService(clientDAO *dao.ClientDAO) *ClientService {
	return &ClientService{DAO: clientDAO}
}

func (s *ClientService) Save(ip string) (uuid.UUID, error) {
	clientObj := &models.Client{
		ID:        uuid.New(),
		IPAddress: ip,
	}

	id, err := s.DAO.Create(clientObj)
	if err != nil {
		return id, err
	}

	return id, nil
}
