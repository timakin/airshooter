package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/timakin/airshooter/config"
	"github.com/timakin/airshooter/constant"

	"strconv"
)

var sharedInstance *gorm.DB

// Singleton
func GetDBInstance() (*gorm.DB, error) {
	if sharedInstance == nil {
		setting := config.Load()
		address := setting.DB.User + ":" + setting.DB.Password + "@" + strconv.Itoa(int(setting.DB.Port)) + "/" + setting.DB.Name + "?charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open("mysql", address)
		if err != nil {
			return nil, errors.Wrap(err, constant.ErrDBConnectionFailed)
		}
		db.LogMode(true)
		sharedInstance = db
	}
	return sharedInstance, nil
}
