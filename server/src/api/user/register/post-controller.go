package register

import (
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

func PostController(context *gin.Context, client *ent.Client) *ent.User {
	params := GetParameters(context, client)

	user := Post(context, client, params)

	return user
}
