package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  //serving static files
  r.Static("/assets", "./assets")

  //html rendering

  r.LoadHTMLGlob("templates/*")

  //uploading Multiple Files
  r.MaxMultipartMemory = 8 << 20 

    r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "File Uploads",
		})
	})

	r.POST("/", func(c *gin.Context){
		file, err := c.FormFile("image")
		if err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": "Failed to Upload Files",
			})
		}


		//save the file with destination

		err = c.SaveUploadedFile(file, "assests/uploads" +file.Filename)

		if err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": "Failed to Save File",
			})
		}

		//Render the Page

		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "assests/uploads" +file.Filename ,
		})
	})
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}