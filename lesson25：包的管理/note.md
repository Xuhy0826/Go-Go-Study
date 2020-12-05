Go两个比较重要的环境变量：
1. `GOROOT`，安装Go的路径
2. `GOPATH`，自定义的开发者的workspace   

在控制台使用命令`go env`可以查看到当前Go定义的环境变量。
Go编译器查找包的顺序：`GOROOT` → `GOPATH`；如果无法找到会引发编译异常。比如引入了`fmt`包，编译器会查找到`C:\Go\src\fmt`。

#### 使用go mod
使用`go mod`命令可以帮助我们生成一个项目模块而摆脱必须在`%GOPATH%`下来创建自己项目的尴尬。创建一个示例项目为例。   
首先`cd`到准备存放项目的目录，这里路径名称为`demo.Pkg`。接着执行`go mod init %项目名%`，按照惯例来说路径名和项目名称应该一致比较好。
```
> go mod init demo.Pkg

go: creating new go.mod: module demo.Pkg
```
可以看到提示，初始化成功了。命令帮我创建了`go.mod`文件，包含这个文件的目录便是模块的根目录了。

#### 引入自定义包


#### 支持远程包导入
很常见的情况是包在GitHub上，如果现在要导入一个远程的包，比如要引用postgres的驱动`import "github.com/bmizerany/pq"`, 编译在导入它时，会先在`GOPATH`下搜索这个包，如果没有，会在使用`go get`命令来获取远程的包，并且会把获取到的源代码存储在GOPATH目录下对应URL的目录里。

