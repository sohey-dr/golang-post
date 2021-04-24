package route

import (
	"strconv"

	"github.com/sohey-dr/golang-post/db"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	posts := db.DbGetAll()
	ctx.HTML(200, "index.html", gin.H{
		"posts": posts,
	})
}

func Create(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	db.DbInsert(text, status)
	ctx.Redirect(302, "/")
}

func Detail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	post := db.DbGetOne(id)
	ctx.HTML(200, "detail.html", gin.H{"post": post})
}

func Update(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	db.DbUpdate(id, text, status)
	ctx.Redirect(302, "/")
}

func DelConf(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	post := db.DbGetOne(id)
	ctx.HTML(200, "delete.html", gin.H{"post": post})
}

func Delete(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	db.DbDelete(id)
	ctx.Redirect(302, "/")

}

func Login(ctx *gin.Context) {
	ctx.HTML(200, "login.html", gin.H{})
}

func Signup(ctx *gin.Context) {
	ctx.HTML(200, "signup.html", gin.H{})
}
