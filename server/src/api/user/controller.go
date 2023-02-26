package user

import (
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, client *ent.Client) {
	r.GET("/users", func(c *gin.Context) {
		users := GetService(client, c)
		c.JSON(200, users)
	})
}
