package register

import (
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

func Post(context *gin.Context, client *ent.Client, params Params) *ent.User {
	user, error := RegisterUser(client, context, params)
	if error != nil {
		apiError := UserRegisterationError()
		context.JSON(apiError.Status, apiError)
	}

	return user
}
