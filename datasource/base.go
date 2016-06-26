package datasource

import (
	//	"github.com/timakin/airshooter/config"
	"github.com/timakin/airshooter/constant"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var sharedInstance *gorm.DB

// Singleton
func GetDBInstance() (*gorm.DB, error) {
	if sharedInstance == nil {
		db, err := gorm.Open("mysql", "root:@/airshooter?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			return nil, errors.Wrap(err, constant.ErrDBConnectionFailed)
		}
		sharedInstance = db
	}
	return sharedInstance, nil
}
