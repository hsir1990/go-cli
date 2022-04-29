 go get -u github.com/kataras/iris
 //被称为速度最快的go后端开发框架


//特点----
专注于高性能
健壮的静态路由支持和通配符子域名支持
视图系统支持超过5以上模板，完全兼容 html/template
支持定制事件的高可扩展性Websocket API,其API类似于socket.io
带有GC, 内存 & redis 提供支持的会话
强大的中间件和插件生态系统
完整 REST API
能定制 HTTP 错误
源码改变后自动加载(热重启)

http1.0只有get post head
http1.1的时候增加的 option put  delete trace connect
一共有8种


put
可以接收xml请求

然后现在最新版本是12的iris，低版本的不建议使用了，变动太大