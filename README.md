# altgotech demo

实现简单个人博客系统的API，不限语言和框架，包含以下功能

1. 用户注册和登录

2. 用户可添加版块或者加入版块

3. 发表文章

4. 评论和点赞

5. 可设置版块管理员，管理员有置顶的权限

6. 全文搜索功能


## 简介

项目中使用到 `mysql` 和 `redis` 请在启动前确保 能正常连接的 `mysql` 和 `redis`，相关的配置文件在 [config/config.json](config/config.json)

--config 可以指定配置文件的路径，不指定默认 [config/config.json](config/config.json)

```bash
go run main.go
```

[Postman测试地址](https://www.postman.com/orange-sunset-547788/workspace/demo/collection/3972752-a7e0121c-18c7-4245-b1ab-b85a1ef34631?action=share&creator=3972752)

[或者Postman接口的json](./demo.postman_collection.json)

