/*
 * Copyright (c) 2019
 * Created at 7/15/19 6:50 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package services

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/oktopriima/privy-id/domain/model"
	"github.com/oktopriima/privy-id/domain/repository"
)

type productCategoryServices struct {
	db *gorm.DB
}

func NewProductCategoryServices(db *gorm.DB) repository.ProductCategoryRepository {
	return &productCategoryServices{db}
}

func (srv *productCategoryServices) Create(category *model.ProductCategory, tx *gorm.DB) (m *model.ProductCategory, err error) {
	db := tx.Create(&category)
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

func (srv *productCategoryServices) FindBy(criteria map[string]interface{}) ([]*model.ProductCategory, error) {
	var arr []*model.ProductCategory

	rows, err := srv.db.Model(&model.ProductCategory{}).Where(criteria).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		m := new(model.ProductCategory)
		err := rows.Scan(&m.ID, &m.ProductID, &m.CategoryID)
		if err != nil {
			return nil, err
		}

		arr = append(arr, m)
	}

	return arr, nil
}

func (srv *productCategoryServices) Update(category *model.ProductCategory, tx *gorm.DB) (err error) {
	err = tx.Save(&category).Error
	return
}

func (srv *productCategoryServices) Delete(category *model.ProductCategory, tx *gorm.DB) (err error) {
	err = tx.Delete(&category).Error
	return
}

func (srv *productCategoryServices) FindOneBy(criteria map[string]interface{}) (*model.ProductCategory, error) {
	m := new(model.ProductCategory)
	row := srv.db.Model(&m).Where(criteria).Scan(&m)
	if err := row.Error; err != nil {
		return nil, err
	}
	return m, nil
}
