package utils

/****************************************
@Author : SleepWalker
@Description:
@Time : 2024/1/3 12:51
****************************************/

func ToAlphaString(index int) string {
	// 定义大写字母A的ASCII码
	const (
		uppercaseA = 'A'
		// 字母表长度
		alphabet = 26
	)

	// 定义结果字符串
	var result []byte
	// 当index大于0时，循环执行
	for index > 0 {
		// index减1
		index--
		// 将index除以alphabet的余数加上大写字母A的ASCII码，添加到结果字符串中
		result = append([]byte{byte(index%alphabet) + uppercaseA}, result...)
		// 将index除以alphabet
		index /= alphabet
	}
	// 返回结果字符串
	return string(result)
}
