# aikucun-sdk-go

爱库存软件开发工具包(Go)

## 接口文档

- [接口文档](http://doc.delivery.aikucun.com/web/#/22?page_id=454)

## 单元测试

鉴于安全原因, 不便在代码直接暴露 appId/appSecret 等信息

所以在测试命令后提供了如下参数

```
go test -v -args -appId xxx -appSecret xxx -erp xxx -erpVersion xxx
```
