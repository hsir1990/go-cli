bee api Apiproject 是只提供接口的项目


bee run  执行命令

import (
	_ "Newproject/routers"

	beego "github.com/beego/beego/v2/server/web"
)
这个下滑线 _ 说明这个包里有一个init 方法可以执行，其余的都不能执行

beego有热更新，修改代码自动更改部署

beego.run()
//beego.run()第一步就是解析配置文件，加载配置内容
1）解析配置文件，也就是我们的app.conf文件，比如端口，应用名称等信息。
2）检查是否开启session，如果开启session，就会初始化一个session对象。
3）是否编译模板，beego框架会在项目启动的时候根据配置把views目录下的所有模板进行预编译，然后存放在map中，这样可以有效的提高模板运行的效率，不需要进行多次编译。
4）监听服务端口。根据app.conf文件中的端口配置，启动监听。