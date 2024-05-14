package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 让每一个controller下面能定义同名函数，使用结构体

type UserController struct{}

func (u UserController) DownloadFile(c *gin.Context) {
	tid, _ := strconv.Atoi(c.Query("id")) // str to int

	if tid == 1 {
		// 读取文件data.json的内容
		filePath := "./static/data.json"   // 经测试，是相对于main.go的
		file, err := os.Open(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		// 获取当前时间戳
		timestamp := time.Now().Unix()

		// 构造文件名，以时间戳为名称
		filename := strconv.FormatInt(timestamp, 10) + ".json"

		// 设置响应头，指定文件名为b.json   // setting same to python3
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Data(http.StatusOK, "application/json", nil)
		// 将文件内容写入响应体
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if tid == 7 {
		// read
		filePath := "./static/data.7z"
		file, err := os.Open(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		// 获取当前时间戳
		timestamp := time.Now().Unix()

		// 构造文件名，以时间戳为名称
		filename := strconv.FormatInt(timestamp, 10) + ".7z"

		// 设置响应头，指定文件名为data.7z   // setting same to python3
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Data(http.StatusOK, "application/x-7z-compressed", nil)

		// 将文件内容写入响应体
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

}

func (u UserController) UploadFile(c *gin.Context) {
	// parse upload file
	file, err := c.FormFile("file") // post form data key: file
	if err != nil {
		// handle error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// handle upload file
	// generate file path
	filePath := "./upload/" + file.Filename

	// save file to local specific path
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		// handle save file error
		c.JSON(http.StatusInternalServerError, gin.H{"erorr": err.Error()})
		return
	}

	// save file success
	// 处理逻辑
	c.JSON(http.StatusOK, gin.H{
		"msg":       "upload file success",
		"file_name": file.Filename,
		"file_path": filePath,
	})
}
