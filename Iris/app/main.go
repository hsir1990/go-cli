package main

import (
	"github.com/kataras/iris"
	iris "github.com/kataras/iris/v12"
)

type person struct {
	name string
}

func main() {
	// //1创建一个app结构体对象，源码中New出一个结构体对象
	// app := iris.New()

	// //2端口监听  //自定义端口    //如果没有这个端口回提示错误
	// app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))

	// // //也可以这么写
	// // app.Run(iris.Addr(":7999"))

	// app.Get("/getRequest", func(context iris.Context) {
	// 	fmt.Println("11")
	// 	//处理get请求，请求的url为 /getRequest
	// 	path := context.Path()
	// 	app.Logger().Info(path) //打印到控制台

	// context.WriteString("请求路径："+path) //发送到浏览器
	//获取get请求的参数
	//username := context.URLParam("username")
	// app.Logger().Info(username)
	// context.HTML("<h1>"+username+"</h1>")  //用html的形式输出到浏览器上
	// })

	//post请求
	// app.Post("/getRequest", func(context iris.Context) {
	// 	fmt.Println("11")
	// 	//处理get请求，请求的url为 /getRequest
	// 	path := context.Path()
	// 	app.Logger().Info(path) //打印到控制台

	// context.WriteString("请求路径："+path) //发送到浏览器
	//获取get请求的参数
	//username := context.PostValue("username")
	// app.Logger().Info(username)
	// context.HTML("<h1>"+username+"</h1>")  //用html的形式输出到浏览器上
	// })

	// //post获取json请求
	// app.Post("/postJson", func(context iris.Context) {
	// 	var person person
	// 	if err := context.ReadJson(&person); err != nil { //获取person的json对象   //ReadXML可以读取xml格式，并转成struct对象
	// 		panic(err.Error())
	// 	}
	// 	//写到前端
	// 	context.WriteF("Received: %#+v \n", person)
	// context.JSON(iris.Map{"str":"aa"})  //用json的格式传到前端
	// })

	// //还有一种handler的写法

	// app.Handle("GET", "/userinfo", func (context iris.Context)  {
	// 	path := context.Path()
	// 	app.Logger().Info(path)

	// })

	// //正则表达式获取路径
	//{login:bool}只能是true或者false
	// app.Get("/weather/{data}/{login:bool}", func(context, iris.Context) {
	// 	path := context.Path()
	// 	data := context.Params.Get("data")
	// if login {
	// 	fmt.Println("已登陆")
	// }else{
	// 	fmt.Println("不登陆")
	// }
	// })

	//新版   见https://www.iris-go.com/docs/#/?id=using-get-post-put-patch-delete-and-options
	app := iris.New()

	booksAPI := app.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)

	}

	app.Listen(":8000")
	type Book struct {
		Title string `json:"title"`
	}

}

type bk1 interface {
}
type Book struct {
	// interface{}
	bk1
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
