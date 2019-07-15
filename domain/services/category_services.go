/*
 * Copyright (c) 2019
 * Created at 7/15/19 5:40 PM
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

type categoryServices struct {
	db *gorm.DB
}

func NewCategoryServices(db *gorm.DB) repository.CategoryRepository {
	return &categoryServices{db}
}

func (srv *categoryServices) Create(category *model.Category, tx *gorm.DB) (m *model.Category, err error) {
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

func (srv *categoryServices) Find(ID int64) (*model.Category, error) {
	m := new(model.Category)
	m.ID = ID

	row := srv.db.Find(&m).Scan(&m)
	if err := row.Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (srv *categoryServices) FindAll() ([]*model.Category, error) {
	var arr []*model.Category

	rows, err := srv.db.Model(&model.Category{}).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		m := new(model.Category)
		err := rows.Scan(&m.ID, &m.Name, &m.Enable)
		if err != nil {
			return nil, err
		}
		arr = append(arr, m)
	}
	return arr, nil
}

func (srv *categoryServices) Update(category *model.Category, tx *gorm.DB) (err error) {
	err = tx.Save(&category).Error
	return
}

func (srv *categoryServices) Delete(category *model.Category, tx *gorm.DB) (err error) {
	err = tx.Delete(&category).Error
	return
}
