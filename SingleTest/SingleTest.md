# 单元测试相关

## 基本命令相关
* 测试包下全部测试用例
```
  go test -v
```
* 指定包下某个单侧用例
```
  go test -run=${匹配函数名} -v
```
* 跳过某些用例
```
  go test -short
```
```
  func TestTimeConsuming(t *testing.T) {
      if testing.Short() {
          t.Skip("short模式下会跳过该测试用例")
      }
  }
```
* 自动生成单测
```
  go get -u github.com/cweill/gotests/...
```
* 测试覆盖率
    * 查看测试覆盖率
  ```
    go test -cover
  ```
    * 覆盖率相关的记录信息输出到一个文件
  ```
    go test -cover -coverprofile=c.out
  ```  
    * 生成html报告信息
  ```
    go tool cover -html=c.out
  ```
* assert断言
```
  go get github.com/stretchr/testify
```