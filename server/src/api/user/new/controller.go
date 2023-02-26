package new

import (
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, client *ent.Client) {
	r.POST("/user/new/:id", func(context *gin.Context) {
		user := PostController(context, client)
		context.JSON(201, user)
	})
}
