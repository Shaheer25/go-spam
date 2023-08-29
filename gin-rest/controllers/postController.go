package controllers

import (
	"github.com/Shaheer_25/initializers"
	model "github.com/Shaheer_25/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct{
		Title string
		Body string
	}

	c.Bind((&body))
	post := model.Post{Title: body.Title, Body:body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error!=nil{
		c.Status(400)
		return
	}


	c.JSON(200, gin.H{
		"Post": post,
	})

}

func PostShow(c *gin.Context){
	var posts []model.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func SinglePostShow (c *gin.Context){
	
	id:=c.Param("id")

	var post model.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}


func PostsUpdate(c *gin.Context){
	
	
	id:=c.Param("id")


	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)


	var post model.Post
	initializers.DB.First(&post, id)


	initializers.DB.Model(&post).Updates(model.Post{
		Title: body.Title ,
		Body: body.Body,
	})


	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context){
	id:=c.Param("id")


	var post model.Post
	initializers.DB.Delete(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}