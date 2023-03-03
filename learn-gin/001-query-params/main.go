package main

import (
	"github.com/gin-gonic/gin"
)

type User struct{
	Name string	`form:"name"`
	Surname string	`form:"surname"`
}

// localhost:5000/users?name=fatih&surname=şen
func Users(c *gin.Context) {
	var b User
    c.Bind(&b)
    c.JSON(200, gin.H{
        "name": b.Name,
        "surname": b.Surname,
    })
}

func main() {
   r := gin.Default()
   r.GET("/users", Users)

   r.Run(":5000") // localhost:5000
}