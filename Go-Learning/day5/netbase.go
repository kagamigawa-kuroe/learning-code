package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler有一点很奇怪的地方在于
//如果你打开ResponseWriter这个类
//你会发现他是个接口
//为什么能对接口进行直接操作呢

//这就要涉及更底层的代码
//在server.go中 我们可以看到如下代码
//func (c *conn) readRequest(ctx context.Context) (w *response, err error) {
//这是一个原型函数
//可以看到里面参数为 *response
//也就是说 我们在hander中传入的接口
//其实会在底层被更具体地实现
//也就是response类
//所哟handler的本质其实是
//func handler(w *http.Response, r *http.Request) {
//而reponse中 实现了ResponseWriter中的write等一系列函数
//所以我们可以认为其是实现了多态

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	//http库中的HandleFunc函数 调用了ServeMux类的HandleFunc函数
	//HandleFunc中又调用了Handle这个函数
	//Handle函数如下
	//mux.Handle(pattern, HandlerFunc(handler))
	//其作用是注册了一个parttern 也就是url格式
	//和对应的一个handler函数 就是处理此url的方法
	//这里pattern是/ 也就是所有的路由都会被此hander处理

	//handler函数的格式为 第一个参数是http.ResponseWriter
	//第二个是 http.Request
	//顾名思义
	//第一个是用于向请求发送信息
	//request记录了一些请求的信息
	//如请求方法（get，post...）
	//如url
	//如http版本

	http.HandleFunc("/", handler)

	//log 记录从8080发过来的所有请求
	log.Fatal(http.ListenAndServe(":8080", nil))

	//这里在看了一下午的源码之后 终于大概看懂了
	//go是如何监听并且处理请求的
	//实在是让人头痛

	//首先 http.ListenAndServe函数
	//会创建一个server对象
	//并调用对象的server.ListenAndServe函数
	//该函数是另一个更底层函数Serve的封装
	//serve函数为每个新的连接请求创建了专门的处理
	//也就是用go函数创建了新的协程
	//c := srv.newConn(rw)
	//go c.serve(connCtx)
	//在协程的serve函数中
	//if fn := c.server.TLSNextProto[proto]; fn != nil {
	//	h := initALPNRequest{ctx, tlsConn, serverHandler{c.server}}
	//从server的TLSNextProto中
	//也就是存储之前所有绑定的handler的对象中
	//匹配请求url对应的处理方式
	//然后进行处理

	//值得注意的是 Servemux这个概念
	//他是一个请求处理复用器
	//如果你不创建一个servemux
	//直接使用http.HandleFunc去绑定一个handler
	//底层代码会自动创建一个defaulservemux去存储所有的匹配规则以及handler
	//而在下面用ListenAndServe去监听时
	//如果第二个参数为nil
	//则会采用defaulservemux为默认的处理器

	//当然你也可以自己创建一个servemux
	//然后再绑定handler以及listen的时候都
	//指定这个mux
}
