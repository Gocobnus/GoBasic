package goready

// vender机制 优先根据根目录下的vender寻找依赖
// go module
// GO111MODULE=off禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
// GO111MODULE=on启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
// GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。

// go mod 常用命令
// go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
// go mod edit        编辑go.mod文件
// go mod graph       打印模块依赖图
// go mod init        初始化当前文件夹, 创建go.mod文件
// go mod tidy        增加缺少的module，删除无用的module
// go mod vendor      将依赖复制到vendor下
// go mod verify      校验依赖
// go mod why         解释为什么需要依赖

// replace 可以在go.mod中使用replace替换成github上对应的库
//replace (
//golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
//golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
//golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
//)

// go get
// go get -u 升级到最新的次要版本
// go get -u=patch升级到最新的订阅版本
// go get package@version 升级到指定version

// 导入本地包
// 非相同项目中使用replace相对路径代替
// require "mypackage" v0.0.0
// replace "mypackage" => "../mypackage"
