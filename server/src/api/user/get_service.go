package user

import (
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

// TODO : 登録しているuserがいなくてもhttp errorとはならないので修正する
// NOTE : テスト用にカスタムエラーを呼び出す
func GetService(client *ent.Client, c *gin.Context) []*ent.User {
	users, err := GetRepository(client, c)
	if err != nil {
		apiError := CreateGetUserError()
		c.JSON(apiError.Status, apiError)
		return nil
	}
	return users
}
