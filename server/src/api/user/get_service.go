package user

import (
	"oruka/ent"

	"github.com/gin-gonic/gin"
)

// TODO : 登録しているuserがいなくてもhttp errorとはならないので修正する
// NOTE : テスト用にカスタムエラーを呼び出す
func GetService(client *ent.Client, context *gin.Context) []*ent.User {
	users, err := GetRepository(client, context)
	if err != nil {
		apiError := CreateGetUserError()
		context.JSON(apiError.Status, apiError)
		return nil
	}
	return users
}
