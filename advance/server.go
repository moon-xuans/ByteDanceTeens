package main

import (
	"byteDanceTeens/advance/controller"
	"byteDanceTeens/advance/repository"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init("./advance/data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
