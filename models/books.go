package models

type MBook struct {
	ID     int64  `gorm:"column:id"`
	Name   string `gorm:"column:name"`
	Author string `gorm:"column:author"`
}

func (book *MBook) TableName() string {
	return "m_book"
}
