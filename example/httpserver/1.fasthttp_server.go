//例子：以FastHttp的方式启动Http Server，同时处理3个（可以处理任意个）不同的公众号的请求。当用户发送任意文字的时候，公众号返回："FastHttp：这是文字！！ 你的应用Id："
//1、你需要在配置好公众号后台
//2、执行 go run 1.fasthttp_server.go 开启server
//3、设置你的公众号的接口URL为：域名/callback/你程序设置的AppNo

package main

import (
	"github.com/hansedong/wechat"
)

func main() {
	FastHttpServe()
}

//以fasthttp的方式启动web server对外提供服务
//个人推荐fasthttp形式的web服务，因为fasthttp要比Go内置的net/http性能高很多，当然，不论是fasthttp还是net/http形式的web服务，通常都满足
//性能要求。fasthttp毕竟是第三方的，而Go的net/http是Go语言官方的，可以依个人喜好选择
func FastHttpServe() {

	//1、开启一个 fasthttp server
	server := wechat.NewFastHttpServer()

	//2.1、实例化一个微信公众号应用
	app := wechat.NewApp()
	app.Configure.Token = "你的Token"
	app.Configure.AppId = "你的AppId"
	app.Configure.AppSecret = "你的密钥"
	app.Configure.AppNo = "可自由设置的应用编号" //应用编号，每个应用必须不同（建议和AppId用同一个值）
	//设置消息处理器
	app.AddTextHandler(MyTextHandler)

	//2.2、你还可以再实例化一个微信公众号应用
	app2 := wechat.NewApp()
	app2.Configure.Token = "你的Token"
	app2.Configure.AppId = "你的AppId"
	app2.Configure.AppSecret = "你的密钥"
	app2.Configure.AppNo = "可自由设置的应用编号"
	app2.Configure.EncodingAESKey = "Zd8P9Ba51FhoWH8NXAJQV2Ghhdssa9zVitQdCqRf7H6"
	app2.Configure.EnableMsgCrypt = true //启用消息安全模式（前提是你的公众号消息加解密方式那里，选择的是安全模式）
	//设置消息处理器
	app2.AddTextHandler(MyTextHandler)

	//2.3、再实例化一个微信公众号应用
	app3 := wechat.NewApp()
	app3.Configure.Token = "你的Token"
	app3.Configure.AppId = "你的AppId"
	app3.Configure.AppSecret = "你的密钥"
	app3.Configure.AppNo = "可自由设置的应用编号"
	//设置消息处理器
	app3.AddTextHandler(MyTextHandler)

	//3、实例化一个公众号应用管理器
	appManager := wechat.NewWechatAppManager()
	appManager.AddApp(app)
	appManager.AddApp(app2)
	appManager.AddApp(app3)

	//4、绑定web服务器和应用管理器
	server.HandleWechat(appManager)
	//5、开启web服务并处理请求
	server.Listen(":9898")

	//6、由于我们上面设置为同时处理3个公众号的请求，所以，每个公众号回调地址
	//都需要设置为：http://xxx.com/callback/你的应用编号。
}

//文本消息处理器
//本例为处理用户消息，返回一个文字
func MyTextHandler(ctx *wechat.WechatContext, app *wechat.WechatApp) interface{} {
	res := ctx.GetMsgTextResponse()
	res.Content = "FastHttp：这是文字！！ 你的应用Id：" + app.Configure.AppId
	return res

}

//文本消息处理器
//本例为处理用户消息，返回当前公众号原始id
//func MyTextHandler(ctx *wechat.WechatContext, _ *wechat.WechatApp) interface{} {
//	res := ctx.GetMsgTextResponse()
//	msg, _ := ctx.DecodeTextMsg()
//	res.Content = "你好，" + msg.ToUserName
//	return res
//}
