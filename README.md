# dirTotal

#### 目录结构
> config 配置文件  
> route 路由  
> service 服务  
> test 单元测试  
> utils 工具  
> main.go 项目入口  

#### 启动项目
`
go run main.go -p /Users/admin/www/learn/go/src/dir/
`
#### http方法 get 请求
```
127.0.0.1:8080/dir?path=config # 获取目录详情
127.0.0.1:8080/dir_info?path=config # 统计目录
127.0.0.1:8080/dir_info?path=dir_http # http 请求统计目录
```
