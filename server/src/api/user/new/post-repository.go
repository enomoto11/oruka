package new

import (
	"oruka/ent"
	"oruka/ent/user"

	"github.com/gin-gonic/gin"
)

func GetUser(client *ent.Client, context *gin.Context, params Params) *ent.User {
	user, _ := client.User.Query().Where(user.ManavisID(params.manavis_id)).Only(context)

	return user
}

func RegisterUser(client *ent.Client, context *gin.Context, params Params) (*ent.User, error) {
	user, error := client.User.
		Create().
		SetFirstName(params.first_name).
		SetLastName(params.last_name).
		SetManavisID(params.manavis_id).
		SetGrade(params.grade).
		Save(context)

	return user, error
}
