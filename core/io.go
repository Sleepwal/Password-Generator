package core

import (
	"fmt"
)

// Input 结构用于处理用户输入
type Input struct{}

// GetPasswordName 获取用户输入的密码名称
func (i *Input) GetPasswordName() string {
	var name string
	fmt.Print("请输入密码名称： ")
	fmt.Scanln(&name)
	return name
}

// Output 结构用于处理输出信息
type Output struct{}

// PrintPassword 打印生成的密码
func (o *Output) PrintPassword(password string) {
	fmt.Printf("生成的密码：%s\n", password)
}

// PrintSaveSuccess 打印保存成功信息
func (o *Output) PrintSaveSuccess(filename, name string) {
	fmt.Printf("密码已保存到 %s 文件中，名称为 %s\n", filename, name)
}
