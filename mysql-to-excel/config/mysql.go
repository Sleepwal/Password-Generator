package config

import "fmt"

/****************************************
@Author : SleepWalker
@Description:
@Time : 2024/1/3 12:14
****************************************/

type MySQL struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
}

func NewMySQLConfig() *MySQL {
	return &MySQL{
		Username: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	}
}

func (m *MySQL) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password,
		m.Host, m.Port,
		m.Database)
}
