# password-generator
密码生成器
- 仅字母 alphanumeric
- 字母和数字 mixedcase
- 字母和数字、特殊字符 allchars

```shell
go run main.go -length=16 -name=my_password -mode=alphanumeric
go run main.go -length=16 -name=my_password -mode=mixedcase
go run main.go -length=16 -name=my_password -mode=allchars
```