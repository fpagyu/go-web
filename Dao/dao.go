package dao

import (
	"github.com/jinzhu/gorm"
)

type DaoIface interface {
	DB(...interface{}) *gorm.DB
}

type DaoBase struct {
}

func (d *DaoBase) Model(val interface{}) *gorm.DB {
	v, ok := val.(string)
	if ok {
		return DB.Table(v)
	} else {
		return DB.Model(val)
	}
}

func BuildSelect(db *gorm.DB, cond map[string]interface{}) *gorm.DB {
	if len(cond) == 0 {
		return db
	}

	if v, has := cond["_fields"]; has {
		delete(cond, "_fields")
		db = db.Select(v)
	}

	if v, has := cond["_order"]; has {
		delete(cond, "_order")
		db = db.Order(v)
	}

	if v, has := cond["_offset"]; has {
		delete(cond, "_offset")
		db = db.Offset(v)
	}

	if v, has := cond["_limit"]; has {
		delete(cond, "_limit")
		db = db.Limit(v)
	}

	if len(cond) > 0 {
		db = db.Where(cond)
	}
	return db
}

type SqlExpr struct {
	sql  string
	args []interface{}
}

func DoTransaction(db *gorm.DB, exprs ...*SqlExpr) (err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, expr := range exprs {
			if err := tx.Exec(expr.sql, expr.args...).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return
}
