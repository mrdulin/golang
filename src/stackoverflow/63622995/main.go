package main

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mrdulin/golang/src/stackoverflow/63622995/db"
)

type Order struct {
	order_id string
}

type OrderPersister struct {
	DB db.OrmDB
	//DB *gorm.DB
}

func (p *OrderPersister) GetOrder(id uuid.UUID) (*Order, error) {
	ret := &Order{}

	err := p.DB.Table("orders").Where("order_id = ?", id).Scan(ret).Error
	return ret, err
}
