package user

import (
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

// TODO : testDBを用意し、repositoryのtestも作成する
func GetRepository(client *ent.Client, context *gin.Context) ([]*ent.User, error) {
	users, err := client.User.Query().All(context)
	return users, err
}
