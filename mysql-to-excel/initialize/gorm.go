package initialize

import (
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mysql-to-excel/config"
	"os"
)

/****************************************
@Author : SleepWalker
@Description:
@Time : 2024/1/3 12:08
****************************************/

func InitGorm() *gorm.DB {
	m := loadConfig()

	dsn := m.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("gorm open err: ", err)
	}

	return db
}

// 加载配置文件
func loadConfig() *config.MySQL {
	m := config.NewMySQLConfig()
	c := &config.Config{Mysql: m}

	file, err := os.ReadFile("./config.yaml")
	if err != nil {
		// 文件不存在，创建一个
		_, _ = os.Create("./config.yaml")
		log.Fatalln("请填写配置文件config.yaml后再运行。")
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatalln("Unmarshal yaml err: ", err)
	}
	return c.Mysql
}
