# gin相关


## 请求类型（GET, POST, PUT, DELETE）
* 通过curl或者浏览器访问
```
  curl -X GET http://localhost:8080/gobasic 
```
* get请求有参数名，一般是在url后通过？添加参数名和参数值，通过quary获取,DefaultQuery可以指定默认值
* get请求无参数名，是在路径中直接加上参数值，需要通过param查询