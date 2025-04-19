# go module

## go module 介绍
go module 是 go 1.11 版本引入的，用来管理 go 模块的依赖关系。
从go 1.13版本开始，go module 成为 go 的默认依赖管理方式。
go module 解决了以下问题：
- 版本控制：go module 可以指定依赖的版本，避免了依赖冲突的问题。
- 依赖管理：go module 可以自动下载和更新依赖，避免了手动管理依赖的麻烦。
- 版本回退：go module 可以方便地回退到之前的版本，避免了因为依赖更新导致的代码不兼容的问题。
- 版本升级：go module 可以方便地升级到最新的版本，避免了因为依赖过旧导致的安全问题。
- 版本锁定：go module 可以锁定依赖的版本，避免了因为依赖更新导致的代码不兼容的问题。
- 版本对比：go module 可以方便地对比不同版本之间的差异，避免了因为依赖更新导致的代码不兼容的问题。
- 版本发布：go module 可以方便地发布新的版本，避免了因为依赖更新导致的代码不兼容的问题。
- 版本回滚：go module 可以方便地回滚到之前的版本，避免了因为依赖更新导致的代码不兼容的问题。

## go module 使用
go module 的使用非常简单，只需要在项目根目录下执行以下命令即可：
```bash
go mod init <module_name>
```
其中，`<module_name>` 是模块的名称，一般是项目的路径，例如 `github.com/user/project`。
执行该命令后，会在项目根目录下生成一个 `go.mod` 文件，文件内容如下：
```go
module <module_name>
go 1.16
```
其中，`go 1.16` 表示该模块的最低支持版本为 go 1.16。
在 `go.mod` 文件中，可以指定依赖的版本，例如：
```go
module <module_name>
go 1.16
require (
    github.com/user/dependency v1.0.0
    github.com/user/another_dependency v2.0.0
)
```

增加依赖后，执行以下命令即可下载依赖：
```bash
go mod tidy
```
该命令会自动下载依赖，并更新 `go.mod` 和 `go.sum` 文件。
`go.sum` 文件中记录了依赖的版本和校验和，用于验证依赖的完整性。
`go.mod` 文件中还可以指定替换依赖的版本，例如：
```go
module <module_name>
go 1.16
require (
    github.com/user/dependency v1.0.0
    github.com/user/another_dependency v2.0.0
)
replace github.com/user/dependency => github.com/user/dependency v1.0.1
```
该命令会将 `github.com/user/dependency` 的版本替换为 `v1.0.1`。
在 `go.mod` 文件中还可以指定排除依赖的版本，例如：
```go
module <module_name>
go 1.16
require (
    github.com/user/dependency v1.0.0
    github.com/user/another_dependency v2.0.0
)
exclude github.com/user/dependency v1.0.0
```
该命令会将 `github.com/user/dependency` 的版本排除为 `v1.0.0`。

## go module 相关命令总结
| 命令 | 说明 |
| ---- | ---- |
| go mod init <module_name> | 初始化 go module |
| go mod tidy | 下载依赖 |
| go test | 测试依赖 |
| go test -cover | 测试覆盖率 |
| go test -coverprofile=coverage.out | 测试覆盖率详细信息 |
| go mod vendor | 下载依赖到 vendor 目录 |
| go mod graph | 显示依赖关系 |
| go mod why | 显示依赖的原因 |
| go mod edit | 编辑 go.mod 文件 |
| go mod download | 下载依赖 |
| go mod verify | 验证依赖的完整性 |
| go mod why -m <module_name> | 显示依赖的原因 |


## go module 导入本地包

### 在同一项目下

> 注意：在一个项目（project）下我们是可以定义多个包（package）的。

#### 目录结构
在moduledemo/main.go中调用了mypackage这个包。
```bash
moduledemo
├── go.mod
├── go.sum
├── main.go
└── mypackage
    ├── mypackage.go
    └── mypackage_test.go
```
#### mypackage.go
```go
package mypackage
import "fmt"
func Hello() {
    fmt.Println("Hello, world!")
}
```
#### mypackage_test.go
```go
package mypackage
import "testing"
func TestHello(t *testing.T) {
    Hello()
}
```
#### main.go
```go
package main
import (
    "moduledemo/mypackage"
)
func main() {
    mypackage.Hello()
}
```
#### go.mod
```go
module moduledemo
go 1.16
```
#### go.sum
```go
github.com/stretchr/testify v1.7.0 h1:4c5
+0x1
github.com/stretchr/testify v1.7.0/go.mod h1:4c5
+0x1
```
#### 执行
```bash
go run main.go
```
输出：
```bash
Hello, world!
```
#### 测试
```bash
go test
```
输出：
```bash
PASS
ok      moduledemo/mypackage  0.001s
```
#### 测试覆盖率
```bash
go test -cover
```
输出：
```bash
PASS
coverage: 100.0% of statements
ok      moduledemo/mypackage  0.001s
```
#### 测试覆盖率详细信息
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```
#### 测试覆盖率详细信息
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```



### 不在同一项目下

#### 目录结构
```bash
├── moduledemo
│   ├── go.mod
│   └── main.go
└── mypackage
    ├── go.mod
    └── mypackage.go

```
#### mypackage.go
```go
package mypackage
import "fmt"
func Hello() {
    fmt.Println("Hello, world!")
}
```
#### mypackage/go.mod
```go
module mypackage
go 1.16
```
#### mypackage/main.go
```go
package main
import (
    "mypackage"
)
func main() {
    mypackage.Hello()
}
```
#### moduledemo/go.mod
因为这两个包不在同一个项目路径下，你想要导入本地包，并且这些包也没有发布到远程的github或其他代码仓库地址。这个时候我们就需要在go.mod文件中使用replace指令。

在调用方也就是moduledemo/go.mod中按如下方式指定使用`相对路径`来寻找mypackage这个包。

```go
module moduledemo
go 1.16
require mypackage v0.0.0
replace mypackage => ../mypackage
```
