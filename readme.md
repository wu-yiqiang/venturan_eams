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
export GO_ENV=dev
go run main.go

开发环境启动
export GO_ENV=test
go run main.go

生产环境启动
export GO_ENV=prod
go run main.go

开发环境打包
export GO_ENV=dev
go build -o venturan main.go

测试环境打包
export GO_ENV=test
go build -o venturan main.go

生产环境打包
export GO_ENV=prod
go build -o venturan main.go

