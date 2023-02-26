package user

import (
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, client *ent.Client) {
	r.GET("/users", func(context *gin.Context) {
		users := GetService(client, context)
		context.JSON(200, users)
	})
}
