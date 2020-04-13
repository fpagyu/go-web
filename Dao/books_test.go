package dao

import (
	"go-web/models"
	"testing"
)

func TestBookDao_DB(t *testing.T) {
	var d BookDao

	var book models.MBook

	db := d.DB()
	db.Where("id=1").Find(&book)

	// db = d.DB(&models.MBook{})
	// db.Where("id=2").Find(&book)
}
