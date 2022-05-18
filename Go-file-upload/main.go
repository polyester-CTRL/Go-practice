package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"log"
	"os"
)

func main() {
	engine:= gin.Default()

	// POST
	engine.POST("/upload", func(req *gin.Context) {
		file, header, err := req.Request.FormFile("myFile")
		if err != nil {
			req.String(http.StatusBadRequest, err.Error())
			return
		}
		fileName := header.Filename
		dir, _ := os.Getwd()
		fp, err := os.Create(dir+"/images/"+fileName)
		defer fp.Close()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(fp, file)
		if err != nil {
			log.Fatal(err)
		}
		req.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// GET
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(req *gin.Context) {
		req.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello gin",
		})
	})
	engine.Run(":3000")
}  