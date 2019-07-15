package services

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/model"
	"github.com/oktopriima/privy-id/domain/repository"
)

type productServices struct {
	db *gorm.DB
}

func NewProductServies(db *gorm.DB) repository.ProductRepository {
	return &productServices{db}
}

func (srv *productServices) Create(product *model.Product, tx *gorm.DB) (m *model.Product, err error) {
	db := tx.Create(&product)
	if err = db.Error; err != nil {
		return
	}

	byteData, err := json.Marshal(db.Value)
	if err != nil {
		return
	}
	if err = json.Unmarshal(byteData, &m); err != nil {
		return
	}

	return
}

func (srv *productServices) Find(ID int64) (*model.Product, error) {
	m := new(model.Product)
	m.ID = ID

	row := srv.db.Find(&m).Scan(&m)

	if err := row.Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (srv *productServices) FindAll() ([]*model.Product, error) {
	var arr []*model.Product

	rows, err := srv.db.Model(&model.Product{}).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		m := new(model.Product)
		err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.Enable)
		if err != nil {
			return nil, err
		}

		arr = append(arr, m)
	}

	return arr, nil
}

func (srv *productServices) Update(product *model.Product, tx *gorm.DB) (err error) {
	err = tx.Save(&product).Error
	return
}

func (srv *productServices) Delete(product *model.Product, tx *gorm.DB) (err error) {
	err = tx.Delete(&product).Error
	return
}
