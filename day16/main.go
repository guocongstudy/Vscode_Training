package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)
//中间件
func timeMiddleWare()gin.HandlerFunc{
	return func(c *gin.Context) {
		begin :=time.Now()
		c.Next()
		useTime :=time.Since(begin)
		log.Printf("use time %d ms\n",useTime)
	}
}


//中间件
func timeMiddleWare2()gin.HandlerFunc{
	return func(c *gin.Context) {
		begin :=time.Now()
		c.Next()
		useTime :=time.Since(begin)
		log.Printf("USE TIME %d ms\n",useTime)
	}
}

func main(){
	engine :=gin.Default()
	engine.Use(timeMiddleWare()) //全局中间件
	engine.GET("/",timeMiddleWare2(), func(c *gin.Context) {//局部中间件
		c.String(http.StatusOK,"OK")
	})
	engine.GET("/home", func(c *gin.Context) {
		c.String(http.StatusOK,"OK")
	})
	engine.Run(":5656")
}