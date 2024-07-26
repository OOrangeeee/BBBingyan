# BBBingyan

<<<<<<< Updated upstream
BB死你

// TODO：add readme
=======
> 一个开箱即用的论坛后端架构。

## 项目简介

BBBingyan是一个开箱即用的论坛后端架构，提供了功能如下：

1. 用户邮箱注册
2. 用户邮箱激活
3. 用户登录通知
4. 用户关注通知
5. 用户@通知
6. 文章发布
7. 各种文章检索
8. 文章评论
9. 给评论评论
10. 文章点赞
11. 查看他人信息、文章、点赞文章等等
12. 管理员功能：删除文章、删除评论、删除用户、创建文章标签，管理文章标签等等
13. 利用WebSocket实现了实时通知功能，并且用RabbitMQ进行了优化
14. 利用JWT进行了用户认证
15. 利用CSRF进行了CSRF防护
16. 维护协程池进行异步邮件发送
17. 利用Reddit的热门文章算法进行了文章排序
18. 等等。。。。

## 项目结构

1. api：API接口文档
2. cmd：主调用程序
3. config：配置文件（需将配置文件放在此文件夹下，具体请看下文）
4. docs：测试用文档，无意义。
5. internal：源码。
   1. configs：项目配置功能
   2. controllers：对外开发的顶层接口控制器
   3. mappers：数据库操作
   4. models：数据库模型与信息传递模型
   5. services：业务逻辑，大部分代码都在这里
   6. utils：工具
6. logs：日志文件
7. scripts：脚本
8. test：测试文件
9. web：前端文件

## 部署方法

本项目提供Dockerfile与docker-compose文件，可以直接使用docker-compose进行部署。

但是要注意部署前需要自建配置文件，具体请看下文。

## 配置文件

配置文件需放在config文件夹下，文件名为config.yaml，内容如下：

请仔细阅读其中需要修改的部分

```yaml
server:
  port: 714
  host: 自己的网址
database:
  dataBaseUserName: "postgres"
  dataBasePassword: "postgres.."
  dataBaseIp: 数据库IP，如果用docker-compose的话就是：db
  dataBasePort: "5432"
  dataBaseName: "bbbingyan"
CSRF:
  cookieSecure: false
  cookieHTTPOnly: true
  cookieMaxAge: 7200
email:
  emailHost: "smtp.163.com"
  emailPort: 465
  emailUserName: 你自己的163邮箱
  emailPassword: 你自己的163邮箱密码
  emailFromNickname: "BBBingyan团队"
  emailOfFollow:
    subject: "BBBingyan关注的人发布新文章了！(Follow Notification)"
    body: |
      <!DOCTYPE html>
      <html lang="zh">
      <head>
          <meta charset="UTF-8">
          <title>关注的人发布新文章通知</title>
      </head>
      <body>
          <p>亲爱的 {用户名},</p>
          <p>您好！</p>
          <p>您关注的人发布了新文章。</p>
          <p>文章标题：{文章标题}</p>
          <p>请您点击下面的链接查看详情：</p>
          <p><a href="文章链接">http://localhost:714/passage/{passage-id}</a></p>
          <p>如果您有任何疑问或需要帮助，请随时回复此邮件或联系我们的客户支持团队。</p>
          <p>祝好，</p>
          <p>BBBingyan团队</p>
          <p>联系电话：{联系电话}</p>
          <p>电子邮箱：{电子邮件地址}</p>
          <p>官方网站：<a href="{官方网站}">{官方网站}</a></p>
      </body>
      </html>
  emailOfAt:
    subject: "BBBingyan@提到通知(Mention Notification)"
    body: |
      <!DOCTYPE html>
      <html lang="zh">
      <head>
          <meta charset="UTF-8">
          <title>提到通知</title>
      </head>
      <body>
          <p>亲爱的 {用户名},</p>
          <p>您好！</p>
          <p>您在BBBingyan服务中被提到了。</p>
          <p>提到您的内容如下：</p>
          <p><a href="文章链接">【替换成你的网址】{passage-id}</a></p>
          <p>请您点击上述链接查看详情。</p>
          <p>如果您有任何疑问或需要帮助，请随时回复此邮件或联系我们的客户支持团队。</p>
          <p>祝好，</p>
          <p>BBBingyan团队</p>
          <p>联系电话：{联系电话}</p>
          <p>电子邮箱：{电子邮件地址}</p>
          <p>官方网站：<a href="{官方网站}">{官方网站}</a></p>
      </body>
      </html>
  emailOfLogin:
    timeRange: 5
    subject: "BBBingyan登录通知(Login Notification)"
    body: |
      <!DOCTYPE html>
      <html lang="zh">
      <head>
          <meta charset="UTF-8">
          <title>登录通知</title>
      </head>
      <body>
          <p>亲爱的 {用户名},</p>
          <p>您好！</p>
          <p>您的账户于 {登录时间} 登录了BBBingyan服务。</p>
          <p>如果这是您本人的操作，登录验证码为：{验证码}。</p>
          <p>如果这不是您本人的操作，请立即修改您的密码以保护您的账户安全。</p>
          <p>如果您有任何疑问或需要帮助，请随时回复此邮件或联系我们的客户支持团队。</p>
          <p>祝好，</p>
          <p>BBBingyan团队</p>
          <p>联系电话：{联系电话}</p>
          <p>电子邮箱：{电子邮件地址}</p>
          <p>官方网站：<a href="{官方网站}">{官方网站}</a></p>
      </body>
      </html>
  emailOfRegister:
    timeRange: 5
    subject: "欢迎使用BBBingyan(Welcome to BBBingyan)"
    body: |
      <!DOCTYPE html>
      <html lang="zh">
      <head>
          <meta charset="UTF-8">
          <title>激活您的账户</title>
      </head>
      <body>
          <p>亲爱的 {用户名},</p>
          <p>您好！</p>
          <p>感谢您注册BBBingyan服务。为了确保您的账户能够顺利使用我们的服务，我们需要您完成账户激活步骤。</p>
          <p>请通过点击下面的激活链接来激活您的账户：</p>
          <p><a href="{激活链接}">点击此处激活您的账户</a></p>
          <p>如果在点击链接时遇到任何问题，您可以将以下链接复制并粘贴到您的网络浏览器地址栏中：</p>
          <p>{激活链接}</p>
          <p>完成激活后，您将能够立即开始使用我们的服务。</p>
          <p>如果您有任何疑问或需要帮助，请随时回复此邮件或联系我们的客户支持团队。</p>
          <p>欢迎加入BBBingyan大家庭，我们期待为您提供卓越的服务体验。</p>
          <p>祝好，</p>
          <p>BBBingyan团队</p>
          <p>联系电话：{联系电话}</p>
          <p>电子邮箱：{电子邮件地址}</p>
          <p>官方网站：<a href="{官方网站}">{官方网站}</a></p>
      </body>
      </html>
info:
  contactPhone: 联系电话
  emailAddress: 邮箱
  webSite: 网站地址
jwt:
  jwtSecret: jwt密钥
passage:
  tags: 文章标签，用逗号分隔，可由程序动态管理
admin:
  adminSecret: 管理员密钥
rabbitmq:
  url: RabbitMQ的URL，用docker-compose的话就是：amqp://rabbitmq:rabbitmq..@rabbitmq:5672/

```
>>>>>>> Stashed changes
