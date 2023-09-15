package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"time"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:PenTest123@tcp(192.168.80.129:3306)/crud_list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 解决插入表的时候会自动添加复数的问题，如：user变成users
			SingularTable: true,
		},
	})

	fmt.Println(db)
	fmt.Println(err)

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.  空闲连接池中的最大连接数
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.  数据库最大连接数
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused. 连接可复用最大时间
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10 seconds

	// struct 结构体定义和优化
	type List struct {
		gorm.Model        // 默认会添加主键等字段
		Name       string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
		State      string `gorm:"type:varchar(20); not null" json:"state" binding:"required"`
		Phone      string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
		Email      string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
		Address    string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
	}

	/*注意点：
	1.结构体中的变量名必须是首字母大写
	gorm    指定类型
	json    表示json接受时的名称
	binding required  表示必传
	*/

	// just update dao use
	////  AutoMigrate 会创建表、缺失的外键、约束、列和索引。 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。 出于保护您数据的目的，它 不会 删除未使用的列
	//db.AutoMigrate(&List{})
	//// https://gorm.io/zh_CN/docs/migration.html

	//// 创建表时添加后缀
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&List{})

	//// 返回当前使用的数据库名
	//cur_name := db.Migrator().CurrentDatabase()
	//fmt.Println(cur_name) // crud_list

	// 1.主键没有（不符合mysql数据库规范）   // 给结构体添加 gorm.Model
	// 2.table name名称变成复数的问题

	// write interface
	PORT := "3001"
	r := gin.Default()

	// TEST
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			//"message": "pong",
			"message": "request success",
		})
	})

	/*业务代码约定
	request success: 200
	request failed: 400
	*/
	// add data
	// ref: https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/
	r.POST("/user/add", func(c *gin.Context) {
		// 定义变量指向结构体
		var data List
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "add failed",
				"data": err.Error(),
				"code": http.StatusBadRequest,
			})
			return
		} else {
			// database operation
			// https://gorm.io/zh_CN/docs/create.html

			result := db.Create(&data) // 通过数据的指针来创建

			fmt.Println(result.Error)
			fmt.Println(result.RowsAffected)

			c.JSON(http.StatusOK, gin.H{
				"msg":  "request success",
				"data": gin.H{},
				//"data": data,
				"code": http.StatusOK,
			})
		}

	})

	// delete data
	/*
		1.找到对应的id所对应的条目
		2.判断id是否存在
		3.如果查到id，则从数据库中删除
		4.如果查不到id，则返回id未找到
	*/
	// restful风格编码规范，推荐用DELETE方法，实际其实也可以用GET或者POST实现。
	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		var data []List

		// receive id
		id := c.Param("id")

		// check whether the id exist
		db.Where("id = ?", id).Find(&data)

		// if id exist, delete it, else report err
		if len(data) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "id is not exist",
				"code": http.StatusNotFound,
				"data": gin.H{},
			})
		} else {
			// operate mysql to delete data item
			// https://gorm.io/zh_CN/docs/delete.html  特别注意：删除一条记录时，删除对象需要指定主键，否则会触发 批量删除

			db.Where("id = ?", id).Delete(&data)

			c.JSON(http.StatusOK, gin.H{
				"msg":  "request success",
				"code": http.StatusOK,
				"data": gin.H{},
			})
		}

	})

	// 另一种非restful风格得实现方式，用get实现删除的示例
	// http://xxx/user/delete?id=123
	r.GET("/user/delete", func(c *gin.Context) {
		id := c.Query("id")
		fmt.Println(id)
	})

	// edit data   // todo 9 https://www.bilibili.com/video/BV1WS4y1t7Py
	r.PUT("/")

	// query data
	r.GET("")

	r.Run(":" + PORT)

}
