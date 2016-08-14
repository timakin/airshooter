package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/timakin/airshooter/config"
	"github.com/timakin/airshooter/constant"

	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
)

var sharedInstance *gorm.DB

// Singleton
func GetDBInstance() (*gorm.DB, error) {
	if sharedInstance == nil {
		setting := config.Load()
		pp.Print(setting.DB)
		address := setting.DB.User + ":" + setting.DB.Password + "@" + strconv.Itoa(int(setting.DB.Port)) + "/" + setting.DB.Name + "?charset=utf8&parseTime=True&loc=Local"
		fmt.Println(address)
		db, err := gorm.Open("mysql", address)
		if err != nil {
			return nil, errors.Wrap(err, constant.ErrDBConnectionFailed)
		}
		db.LogMode(true)
		sharedInstance = db
	}
	return sharedInstance, nil
}

func Init() {
	dbConnection, err := GetDBInstance()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	dbConnection.DropTable("access_tokens")
	dbConnection.DropTable("clients")
	dbConnection.DropTable("message_recipients")
	dbConnection.DropTable("message_senders")
	dbConnection.DropTable("messages")
	dbConnection.DropTable("notification_recipients")
	dbConnection.DropTable("notification_senders")
	dbConnection.DropTable("notifications")

	// Load the raw SQL from resources directory and exec them.
	absSchemaDirPath, _ := filepath.Abs("../resources/schema/")
	sqlfiles, _ := ioutil.ReadDir(absSchemaDirPath)
	reg, _ := regexp.Compile("^create_table_")
	for _, f := range sqlfiles {
		if reg.MatchString(f.Name()) {
			continue
		}
		absSchemaFilePath, _ := filepath.Abs(absSchemaDirPath + f.Name())
		sqlByte, _ := ioutil.ReadFile(absSchemaFilePath)
		rawsql := string(sqlByte)
		dbConnection.Exec(rawsql)
	}
}

func Close() {
	dbConnection, err := GetDBInstance()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	dbConnection.Close()
}
