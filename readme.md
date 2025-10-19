swagger 地址
http://localhost:9527/swagger/index.html

项目初始化
go mod init 项目名
检测依赖
go mod tidy
安装依赖
go mod download
导入依赖
go mod vendor

本地启动
go run main.go --mode dev

开发环境启动
go run main.go --mode test

生产环境启动
go run main.go --mode prod

开发环境打包

测试环境打包

生产环境打包

