package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"
)

func url(engine *gin.Engine) {
	engine.GET("/stu", func(ctx *gin.Context) {
		name := ctx.Query("name")
		addr := ctx.DefaultQuery("addr", "china")
		//ctx.Writer.Write([]byte(name+"live in"+addr))
		//ctx.Writer.WriteHeader(http.StatusOK)

		ctx.String(http.StatusOK, name+"live in"+addr)
	})
}

func form(engine *gin.Engine) {
	engine.POST("/home", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		addr := ctx.DefaultPostForm("addr", "china")
		ctx.String(http.StatusOK, name+"live in"+addr)
	})
}
func restfulform(engine *gin.Engine) {
	engine.POST("/user/:name/:addr", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		addr := ctx.DefaultPostForm("addr", "china")
		ctx.String(http.StatusOK, name+"live in"+addr)
	})
}
func upload_file(engine *gin.Engine) {
	engine.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, "upload file failed")
		} else {
			c.SaveUploadedFile(file, "./date/"+file.Filename)
			c.String(http.StatusOK, file.Filename+"upload success")

		}
	})
}

func upload_files(engine *gin.Engine) {
	engine.POST("/multi_upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for _, file := range files {
			c.SaveUploadedFile(file, "./data/"+file.Filename)
		}
		c.String(http.StatusOK, strconv.Itoa(len(files))+"upload success")
	})
}

type Student struct {
	Name       string    `form:"username" json:"name" uri:"user" xml:"user" yaml:"user" binding:"required"`
	Addr       string    `form:"addr" json:"addr" uri:"addr" xml:"addr" yaml:"addr" binding:"required"`
	Enrollment time.Time `form:"enroll" binding:"required,before_today",time_format:"2006-02-02" time_utc:"8"`
	Graduation time.Time `form:"graduate" binding:"required,before_today",time_format:"2006-02-02" time_utc:"8"`
}

var beforeToday validator.Func = func(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Before(today) {
			return true
		}
	}
	return false
}

func formBind(engine *gin.Engine) {
	engine.LoadHTMLFiles("static/template.html")
	engine.POST("/stu/form", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBind(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse param failed")
		} else {
			//ctx.String(http.StatusOK, stu.Name+" live in"+stu.Addr)
			ctx.JSON(http.StatusOK, stu)

		}
	})
}
func jsonBind(engine *gin.Engine) {
	engine.POST("/stu/json", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBind(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse param failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in"+stu.Addr)
		}
	})
}

func yamlBind(engine *gin.Engine) {
	engine.POST("/stu/yaml", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBind(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse param failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in"+stu.Addr)
		}
	})
}

func uriBind(engine *gin.Engine) {
	engine.POST("/stu/json", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBind(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse param failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in"+stu.Addr)
		}
	})
}

func main() {
	engine := gin.Default()
	url(engine)
	form(engine)
	restfulform(engine)
	upload_file(engine)
	upload_files(engine)

	formBind(engine)
	jsonBind(engine)
	yamlBind(engine)
	uriBind(engine)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("before_today", beforeToday)
	}
	//根目录
	engine.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"Welcome")
	})

	//重定向
	engine.GET("/not_exists", func(c *gin.Context) {
         c.Redirect(http.StatusMovedPermanently,"http://127.0.0.1")
	})
	engine.Run(":5656")
}
