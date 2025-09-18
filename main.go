package main

import (
	"log"
	ginitem "my-app/modules/item/transport/gin"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatalln("⚠️ Missing DB_DSN environment variable")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("/add", ginitem.CreateItem(db))
			items.GET("/get-all", ginitem.ListItem(db))
			items.GET("/get-by-id/:id", ginitem.GetItem(db))
			items.POST("/edit/:id", ginitem.UpdateItem(db))
			items.POST("/delete/:id", ginitem.DeleteItem(db))

		}
	}

	r.Run(":3000")
}
