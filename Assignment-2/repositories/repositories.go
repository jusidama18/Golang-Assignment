package repositories

import (
	"Assignment-2/repositories/models"

	"github.com/jinzhu/gorm"
)

type OrderRepos interface {
	Create(models.Orders) error
	GetMany(models.Orders) error
	GetOne(models.Orders, int) error
	Update(models.Orders) error
	Delete(models.Orders) error
}

type OrdersDB struct {
	db *gorm.DB
}

func MakeOrderRepo(db *gorm.DB) *OrdersDB {
	return &OrdersDB{db: db}
}

func (r *OrdersDB) Create(order models.Orders) error {
	err := r.db.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrdersDB) GetMany(order models.Orders) error {
	err := r.db.Preload("Items").Find(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrdersDB) GetOne(order models.Orders, ids int) error {
	err := r.db.Preload("Items").First(&order, ids).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrdersDB) Update(order models.Orders, ids int) error {
	err := r.db.Where(`OrderID = ?`, ids).Update(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *OrdersDB) Delete(order models.Orders, ids int) error {
	err := r.db.Delete(&order, ids).Error
	if err != nil {
		return err
	}
	return nil
}
