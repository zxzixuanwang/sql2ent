# sql2ent
该项目提供 sql 语句转化为 `entgo schema` 代码的命令行工具, 以提高工作效率。

[前往学习entgo](https://entgo.io)

## 功能

### 已完成
1. 使用命令行批量转化
2. 支持 mysql
3. 读取数据库，批量生成 `schema` 文件。

### 计划
1. 支持更多的数据库，例如：MariaDB、SQLite、PostgreSQL。
2. 支持外键
3. 等等...

## 快速开始

### 第一步：安装 `sql2ent`
```shell
# Go 1.20 或更高版本
go get -u github.com/zxzixuanwang/sql2ent



### 第二步：运行命令
```shell
sql2ent mysql ddl -src "./sql/*.sql" -dir "./ent/schema"
# 或者
sql2ent mysql ddl -target "root:pass@tcp(localhost:3306)/test" -dir "./ent/schema"

```
说明：
* -src: 输入 sql 路径，可模糊匹配
* -dir: 输出目录，默认 `./ent/schema`
* -mysql_target mysql的连接, `root:pass@tcp(localhost:3306)/test`



## 参与开源

1. 点击 Fork
2. 提交自己的代码到 Fork 的仓库中
3. Pull Request 将自己的代码合并