# bzojChecker

一个监测他人BZOJ代码提交的邪恶爬虫。

只要有人提交代码，其他用户就可以接收到消息。

BZOJ 爆炸/好了 会有贴心的提醒。

经实测如果你使用 outlook 邮箱可能会报错，请使用国内邮箱。

# 食用方法

先 [下载](https://github.com/YuhangQ/bzojChecker/releases) 好对应平台的程序。
然后配置好 config.json ，就可以直接运行了，不需要其他依赖。

```
{
  // 下面三个是邮件服务器的配置
  "smtphost": "地址:端口",
  "username": "用户名",
  "password": "密码",
  // 下面是用户的配置，username 是 OJ 用户名，email 是邮箱，receive 表示是否接收邮件。
  "users": [
    {"username": "user1", "email": "user1@example.com", "receive": true},
    {"username": "user2", "email": "user2@example.com", "receive": true}
  ]
}
```
