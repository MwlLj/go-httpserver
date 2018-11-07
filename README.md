# go-httpserver
对 golang 原生的 http server 的路由功能进行增强

之前一直使用的 github.com/julienschmidt/httprouter http路由
但是对于这个路由器的一点非常不喜欢, 就是不能在 handler 函数中使用创建 handler 的类本身

所以自己重新写了一个, 可以在 handler 中获取 server 指针, 并且可以设置用户额外的参数:

type HandleFunc func(w http.ResponseWriter, r *http.Request, param CUrlParam, server CHttpServer, userdata interface{}) bool

用户参数可以在 NewRouterHandler(userdata interface{}, handle HandleFunc) 中指定
(userdata 参数)

另外这里存在一个返回值, 如果用户返回 nil, 内部将 io.WriteString(w, "interior error")
