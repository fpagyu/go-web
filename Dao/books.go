package dao

import (
	"go-web/models"

	"github.com/jinzhu/gorm"
)

type BookDao struct {
	DaoBase
}

func (d BookDao) DB(val ...interface{}) *gorm.DB {
	if val == nil {
		return DB.Table("m_book")
	}

	return d.DaoBase.Model(val[0])
}

func (d BookDao) FindByID(id int) (*models.MBook, error) {
	var book models.MBook
	err := d.DB().First(&book, "id=?", id).Error

	return &book, err
}
