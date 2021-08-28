```
go mod init example.com/printer
go mod init example.com/hello

go mod edit -replace example.com/greetings=../greetings
go mod tidy

go test
go test -v

go build：会生成可执行文件
go list -f '{{.Target}}'：查询`go install`时，会安装到哪里去
export PATH=$PATH:/path/to/your/install/directory
go env -w GOBIN=/path/to/your/bin

go install
现在就可以在任意目录执行`hello`了
```
