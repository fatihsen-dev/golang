package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Person struct {
   // ID   string `uri:"id" binding:"required,uuid"` // uuid type = 987fbc97-4bed-5078-9f07-9141ba07c9f3
   ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func route(c *gin.Context) {
	var person Person
	log.Println(person)
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}

func main()  {
	router := gin.Default()
	router.Use(cors.Default())
   router.GET("/:name/:id", route)
	router.Run(":5000")
}