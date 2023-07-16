package models

import "github.com/google/uuid"

type Info struct {
	ID       uuid.UUID `db:"id" json:"id" type:"uuid,primary key"`
	Key      []byte    `db:"key" json:"key" type:"bytea, not null"`
	Value    []byte    `db:"value" json:"value" type:"bytea,not null"`
	ClientID uuid.UUID `db:"client_id" json:"client_id" type:"uuid,foreign key"`
	ReadOnly bool      `db:"read_only" json:"read_only" type:"bool,default false"`
}
