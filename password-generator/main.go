package main

import (
	"SleepWalker/PasswordGenerator/core"
	"flag"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// 命令行参数
	length := flag.Int("length", 12, "密码长度")
	name := flag.String("name", "default", "密码名称")
	//filename := flag.String("file", "./password-generator/passwords.yaml", "保存密码的YAML文件")
	filename := flag.String("file", "./passwords.yaml", "保存密码的YAML文件")
	mode := flag.String("mode", "alphanumeric", "密码模式: alphanumeric, mixedcase, allchars")
	flag.Parse()

	// 创建配置实例
	cfg := core.NewConfig(*filename)

	// 读取已有密码
	if err := cfg.Load(); err != nil {
		viper.WriteConfigAs(*filename) // 如果配置文件不存在，则创建一个新文件
	}

	// 创建密码生成器实例
	gen := core.NewGenerator(map[string]string{
		"alphanumeric": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"mixedcase":    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"allchars":     "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>/?",
	})

	// 生成密码
	password, err := gen.Generate(*length, *mode)
	if err != nil {
		panic(err)
	}

	// 打印密码
	output := core.Output{}
	output.PrintPassword(password)

	// 保存新密码
	cfg.Passwords[*name] = password
	if err := cfg.Save(); err != nil {
		log.Fatalln("保存密码失败, ", err)
	}

	// 打印保存成功信息
	output.PrintSaveSuccess(*filename, *name)
}

// go run main.go -length=16 -name=my_password -mode=allchars
// main.exe -length=16 -name=my_password -mode=allchars
