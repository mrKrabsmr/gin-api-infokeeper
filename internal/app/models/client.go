package models

import "github.com/google/uuid"

type Client struct {
	ID        uuid.UUID `db:"id" json:"id" type:"uuid,primary key"`
	IPAddress string    `db:"ip_address" json:"ip_address" type:"inet"`
}
