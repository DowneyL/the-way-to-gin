package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("UpdateOn", time.Now().Unix())
}

func GetArticles(page int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(page).Limit(pageSize).Find(&articles)
	return
}

func GetArticlesTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func ExistArticleById(id int) bool {
	var article Article
	db.Where("id = ?", id).First(&article)

	return article.ID > 0
}

func EditArticle(id int, data interface{}) bool {
	wdb.Model(Article{}).Where("id = ?", id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	wdb.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	wdb.Where("id = ?", id).Delete(Article{})
	return true
}
