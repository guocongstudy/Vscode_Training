package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var(
	authMap sync.Map
)

func genCookieName(c *gin.Context)string {
	return base64.StdEncoding.EncodeToString([]byte(c.Request.RemoteAddr))
}

//func login(engine *gin.Engine){
//	engine.POST("/login",func(c *gin.Context){
//		cookie_name:=genCookieName(c)
//		cookie_value:=strconv.Itoa(time.Now().Nanosecond())
//		authMap.Store(cookie_name,cookie_value)
//		c.SetCookie(
//			cookie_name,
//			cookie_value,
//			600,
//			"/",
//			false,
//			true,
//			)
//		c.String(http.StatusOK,"登录成功")
//	})
func authMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context){
		cookie_name:=genCookieName(c)
		var cookie_value string
		for _,cookie :=range c.Request.Cookies(){
			if cookie.Name==cookie_name{
				cookie_value = cookie.Value
				break
			}
		}
		if len(cookie_value) ==0{
			c.Abort()
		}else{
			if v,ok:=authMap.Load(cookie_name);ok{
				if v.(string)==cookie_value{
					c.Next()
				}else {
					c.Abort()
				}
			}
		}
	}
}

func userCenter(engine *gin.Engine){
	engine.POST("/user_center",authMiddleWare(),func(c *gin.Context){
		c.String(http.StatusOK,"这里是你的个人空间")
	})
}

func main(){
	engine :=gin.Default()
	userCenter(engine)
	engine.Run(":5656")
}