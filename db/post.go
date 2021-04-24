package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	gorm.Model
	Text   string
	Status string
}

//DBに接続
func dbOpen() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbOpen）")
	}
	return db
}

//DB初期化
func DbInit() {
	db := dbOpen()
	db.AutoMigrate(&Post{})
	defer db.Close()
}

//DB追加
func DbInsert(text string, status string) {
	db := dbOpen()
	db.Create(&Post{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func DbUpdate(id int, text string, status string) {
	db := dbOpen()
	var post Post
	db.First(&post, id)
	post.Text = text
	post.Status = status
	db.Save(&post)
	db.Close()
}

//DB削除
func DbDelete(id int) {
	db := dbOpen()
	var post Post
	db.First(&post, id)
	db.Delete(&post)
	db.Close()
}

//DB全取得
func DbGetAll() []Post {
	db := dbOpen()
	var posts []Post
	db.Order("created_at desc").Find(&posts)
	db.Close()
	return posts
}

//DB一つ取得
func DbGetOne(id int) Post {
	db := dbOpen()
	var post Post
	db.First(&post, id)
	db.Close()
	return post
}
