# mysql-to-excel
将MySQL数据表信息导出到Excel，包括：
  - 列名
  - 字段类型
  - 长度
  - 是否可为空
  - 默认值
  - 备注

配置文件
```yaml
mysql:
  username: root
  password: root
  host: 127.0.0.1
  port: 3306
  database: test
```